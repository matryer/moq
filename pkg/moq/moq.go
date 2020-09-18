package moq

import (
	"bytes"
	"errors"
	"fmt"
	"go/build"
	"go/types"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"golang.org/x/tools/go/packages"
)

// Mocker can generate mock structs.
type Mocker struct {
	srcPkg     *packages.Package
	tmpl       *template.Template
	pkgName    string
	pkgPath    string
	fmter      func(src []byte) ([]byte, error)
	stubImpl   bool
	skipEnsure bool

	imports map[string]bool
}

// Config specifies details about how interfaces should be mocked.
// SrcDir is the only field which needs be specified.
type Config struct {
	SrcDir     string
	PkgName    string
	Formatter  string
	StubImpl   bool
	SkipEnsure bool
}

// New makes a new Mocker for the specified package directory.
func New(conf Config) (*Mocker, error) {
	srcPkg, err := pkgInfoFromPath(conf.SrcDir, packages.NeedName|packages.NeedTypes|packages.NeedTypesInfo)
	if err != nil {
		return nil, fmt.Errorf("couldn't load source package: %s", err)
	}

	pkgName := conf.PkgName
	if pkgName == "" {
		pkgName = srcPkg.Name
	}

	pkgPath, err := findPkgPath(conf.PkgName, srcPkg)
	if err != nil {
		return nil, fmt.Errorf("couldn't load mock package: %s", err)
	}

	tmpl, err := template.New("moq").Funcs(templateFuncs).Parse(moqTemplate)
	if err != nil {
		return nil, err
	}

	fmter := gofmt
	switch conf.Formatter {
	case "goimports":
		fmter = goimports
	case "noop":
		fmter = noopFmt
	}

	return &Mocker{
		tmpl:       tmpl,
		srcPkg:     srcPkg,
		pkgName:    pkgName,
		pkgPath:    pkgPath,
		fmter:      fmter,
		stubImpl:   conf.StubImpl,
		skipEnsure: conf.SkipEnsure,
		imports:    make(map[string]bool),
	}, nil
}

func findPkgPath(pkgInputVal string, srcPkg *packages.Package) (string, error) {
	if pkgInputVal == "" {
		return srcPkg.PkgPath, nil
	}
	if pkgInDir(".", pkgInputVal) {
		return ".", nil
	}
	if pkgInDir(srcPkg.PkgPath, pkgInputVal) {
		return srcPkg.PkgPath, nil
	}
	subdirectoryPath := filepath.Join(srcPkg.PkgPath, pkgInputVal)
	if pkgInDir(subdirectoryPath, pkgInputVal) {
		return subdirectoryPath, nil
	}
	return "", nil
}

func pkgInDir(pkgName, dir string) bool {
	currentPkg, err := pkgInfoFromPath(dir, packages.NeedName)
	if err != nil {
		return false
	}
	return currentPkg.Name == pkgName || currentPkg.Name+"_test" == pkgName
}

// Mock generates a mock for the specified interface name.
func (m *Mocker) Mock(w io.Writer, names ...string) error {
	if len(names) == 0 {
		return errors.New("must specify one interface")
	}

	doc := doc{
		PackageName: m.pkgName,
		Imports:     moqImports,
		StubImpl:    m.stubImpl,
		SkipEnsure:  m.skipEnsure,
	}

	mocksMethods := false

	tpkg := m.srcPkg.Types
	for _, name := range names {
		n, mockName := parseInterfaceName(name)
		iface := tpkg.Scope().Lookup(n)
		if iface == nil {
			return fmt.Errorf("cannot find interface %s", n)
		}
		if !types.IsInterface(iface.Type()) {
			return fmt.Errorf("%s (%s) not an interface", n, iface.Type().String())
		}
		iiface := iface.Type().Underlying().(*types.Interface).Complete()
		obj := obj{
			InterfaceName: n,
			MockName:      mockName,
		}
		for i := 0; i < iiface.NumMethods(); i++ {
			mocksMethods = true
			meth := iiface.Method(i)
			sig := meth.Type().(*types.Signature)
			method := &method{
				Name: meth.Name(),
			}
			obj.Methods = append(obj.Methods, method)
			method.Params, method.Returns = m.extractArgs(sig)
		}
		doc.Objects = append(doc.Objects, obj)
	}

	if mocksMethods {
		doc.Imports = append(doc.Imports, "sync")
	}

	for pkgToImport := range m.imports {
		doc.Imports = append(doc.Imports, stripVendorPath(pkgToImport))
	}

	if tpkg.Name() != m.pkgName {
		doc.SourcePackagePrefix = tpkg.Name() + "."
		doc.Imports = append(doc.Imports, stripVendorPath(tpkg.Path()))
	}

	var buf bytes.Buffer
	err := m.tmpl.Execute(&buf, doc)
	if err != nil {
		return err
	}
	formatted, err := m.fmter(buf.Bytes())
	if err != nil {
		return err
	}
	if _, err := w.Write(formatted); err != nil {
		return err
	}
	return nil
}

func (m *Mocker) packageQualifier(pkg *types.Package) string {
	if m.pkgPath != "" && m.pkgPath == pkg.Path() {
		return ""
	}
	path := pkg.Path()
	if pkg.Path() == "." {
		wd, err := os.Getwd()
		if err == nil {
			path = stripGopath(wd)
		}
	}
	m.imports[path] = true
	return pkg.Name()
}

func (m *Mocker) extractArgs(sig *types.Signature) (params, results []*param) {
	pp := sig.Params()
	for i := 0; i < pp.Len(); i++ {
		p := m.buildParam(pp.At(i), "in"+strconv.Itoa(i+1))
		// check for final variadic argument
		p.Variadic = sig.Variadic() && i == pp.Len()-1 && p.Type[0:2] == "[]"
		params = append(params, p)
	}

	rr := sig.Results()
	for i := 0; i < rr.Len(); i++ {
		results = append(results, m.buildParam(rr.At(i), "out"+strconv.Itoa(i+1)))
	}

	return
}

func (m *Mocker) buildParam(v *types.Var, fallbackName string) *param {
	name := v.Name()
	if name == "" || name == "_" {
		name = fallbackName
	}
	typ := types.TypeString(v.Type(), m.packageQualifier)
	return &param{Name: name, Type: typ}
}

func pkgInfoFromPath(srcDir string, mode packages.LoadMode) (*packages.Package, error) {
	pkgs, err := packages.Load(&packages.Config{
		Mode: mode,
		Dir:  srcDir,
	})
	if err != nil {
		return nil, err
	}
	if len(pkgs) == 0 {
		return nil, errors.New("No packages found")
	}
	if len(pkgs) > 1 {
		return nil, errors.New("More than one package was found")
	}
	if errs := pkgs[0].Errors; len(errs) != 0 {
		if len(errs) == 1 {
			return nil, errs[0]
		}
		return nil, fmt.Errorf("%s (and %d more errors)", errs[0], len(errs)-1)
	}
	return pkgs[0], nil
}

func parseInterfaceName(name string) (ifaceName, mockName string) {
	parts := strings.SplitN(name, ":", 2)
	ifaceName = parts[0]
	mockName = ifaceName + "Mock"
	if len(parts) == 2 {
		mockName = parts[1]
	}
	return
}

type doc struct {
	PackageName         string
	SourcePackagePrefix string
	Objects             []obj
	Imports             []string
	StubImpl            bool
	SkipEnsure          bool
}

type obj struct {
	InterfaceName string
	MockName      string
	Methods       []*method
}
type method struct {
	Name    string
	Params  []*param
	Returns []*param
}

func (m *method) Arglist() string {
	params := make([]string, len(m.Params))
	for i, p := range m.Params {
		params[i] = p.String()
	}
	return strings.Join(params, ", ")
}

func (m *method) ArgCallList() string {
	params := make([]string, len(m.Params))
	for i, p := range m.Params {
		params[i] = p.CallName()
	}
	return strings.Join(params, ", ")
}

func (m *method) ReturnArgTypeList() string {
	params := make([]string, len(m.Returns))
	for i, p := range m.Returns {
		params[i] = p.TypeString()
	}
	if len(m.Returns) > 1 {
		return fmt.Sprintf("(%s)", strings.Join(params, ", "))
	}
	return strings.Join(params, ", ")
}

func (m *method) ReturnArgNameList() string {
	params := make([]string, len(m.Returns))
	for i, p := range m.Returns {
		params[i] = p.Name
	}
	return strings.Join(params, ", ")
}

type param struct {
	Name     string
	Type     string
	Variadic bool
}

func (p param) String() string {
	return fmt.Sprintf("%s %s", p.Name, p.TypeString())
}

func (p param) CallName() string {
	if p.Variadic {
		return p.Name + "..."
	}
	return p.Name
}

func (p param) TypeString() string {
	if p.Variadic {
		return "..." + p.Type[2:]
	}
	return p.Type
}

var templateFuncs = template.FuncMap{
	"Exported": func(s string) string {
		if s == "" {
			return ""
		}
		for _, initialism := range golintInitialisms {
			if strings.ToUpper(s) == initialism {
				return initialism
			}
		}
		return strings.ToUpper(s[0:1]) + s[1:]
	},
}

// stripVendorPath strips the vendor dir prefix from a package path.
// For example we might encounter an absolute path like
// github.com/foo/bar/vendor/github.com/pkg/errors which is resolved
// to github.com/pkg/errors.
func stripVendorPath(p string) string {
	parts := strings.Split(p, "/vendor/")
	if len(parts) == 1 {
		return p
	}
	return strings.TrimLeft(path.Join(parts[1:]...), "/")
}

// stripGopath takes the directory to a package and removes the
// $GOPATH/src path to get the canonical package name.
func stripGopath(p string) string {
	for _, srcDir := range build.Default.SrcDirs() {
		rel, err := filepath.Rel(srcDir, p)
		if err != nil || strings.HasPrefix(rel, "..") {
			continue
		}
		return filepath.ToSlash(rel)
	}
	return p
}
