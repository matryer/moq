package moq

import (
	"bytes"
	"errors"
	"fmt"
	"go/build"
	"go/format"
	"go/types"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/tools/go/packages"
)

// This list comes from the golint codebase. Golint will complain about any of
// these being mixed-case, like "Id" instead of "ID".
var golintInitialisms = []string{
	"ACL",
	"API",
	"ASCII",
	"CPU",
	"CSS",
	"DNS",
	"EOF",
	"GUID",
	"HTML",
	"HTTP",
	"HTTPS",
	"ID",
	"IP",
	"JSON",
	"LHS",
	"QPS",
	"RAM",
	"RHS",
	"RPC",
	"SLA",
	"SMTP",
	"SQL",
	"SSH",
	"TCP",
	"TLS",
	"TTL",
	"UDP",
	"UI",
	"UID",
	"UUID",
	"URI",
	"URL",
	"UTF8",
	"VM",
	"XML",
	"XMPP",
	"XSRF",
	"XSS",
}

// Mocker can generate mock structs.
type Mocker struct {
	srcPkg  *packages.Package
	tmpl    *template.Template
	pkgName string
	pkgPath string

	aliases map[string]string // path -> name
	imports map[string]string // name -> path
}

// New makes a new Mocker for the specified package directory.
func New(src, packageName string) (*Mocker, error) {
	srcPkg, err := pkgInfoFromPath(src, packages.NeedName|packages.NeedSyntax|packages.NeedTypes|packages.NeedTypesInfo)
	if err != nil {
		return nil, fmt.Errorf("couldn't load source package: %s", err)
	}
	pkgPath, err := findPkgPath(packageName, srcPkg)
	if err != nil {
		return nil, fmt.Errorf("couldn't load mock package: %s", err)
	}

	tmpl, err := template.New("moq").Funcs(templateFuncs).Parse(moqTemplate)
	if err != nil {
		return nil, err
	}
	return &Mocker{
		tmpl:    tmpl,
		srcPkg:  srcPkg,
		pkgName: preventZeroStr(packageName, srcPkg.Name),
		pkgPath: pkgPath,
		aliases: extractAliases(srcPkg),
		imports: make(map[string]string),
	}, nil
}

func preventZeroStr(val, defaultVal string) string {
	if val == "" {
		return defaultVal
	}
	return val
}

func extractAliases(pkg *packages.Package) map[string]string {
	aliases := make(map[string]string)
	for _, syntax := range pkg.Syntax {
		for _, imprt := range syntax.Imports {
			if imprt.Name != nil {
				aliases[strings.Trim(imprt.Path.Value, `"`)] = imprt.Name.Name
			}
		}
	}
	return aliases
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
	}

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
			m.imports["sync"] = "sync"
			meth := iiface.Method(i)
			sig := meth.Type().(*types.Signature)
			method := &method{
				Name: meth.Name(),
			}
			obj.Methods = append(obj.Methods, method)
			method.Params = m.extractArgs(sig, sig.Params(), "in%d")
			method.Returns = m.extractArgs(sig, sig.Results(), "out%d")
		}
		doc.Objects = append(doc.Objects, obj)
	}

	for pkgName, pkgToImport := range m.imports {
		doc.Imports = append(doc.Imports, makeImportLine(pkgName, pkgToImport))
	}

	if tpkg.Name() != m.pkgName {
		doc.SourcePackagePrefix = tpkg.Name() + "."
		doc.Imports = append(doc.Imports, `"`+stripVendorPath(tpkg.Path())+`"`)
	}

	var buf bytes.Buffer
	err := m.tmpl.Execute(&buf, doc)
	if err != nil {
		return err
	}
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("go/format: %s", err)
	}
	if _, err := w.Write(formatted); err != nil {
		return err
	}
	return nil
}

func makeImportLine(pkgName, pkgToImport string) string {
	if strings.HasSuffix(pkgToImport, pkgName) {
		return `"` + stripVendorPath(pkgToImport) + `"`
	}
	return fmt.Sprintf(`%s "%s"`, pkgName, stripVendorPath(pkgToImport))
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
	name := m.aliases[path]
	if name == "" {
		name = pkg.Name()
	}
	for i := 1; ; i++ {
		if p, exists := m.imports[name]; exists && p != path {
			name = fmt.Sprintf("%s%d", pkg.Name(), i)
		} else {
			break
		}
	}
	m.imports[name] = path
	return name
}

func (m *Mocker) extractArgs(sig *types.Signature, list *types.Tuple, nameFormat string) []*param {
	var params []*param
	listLen := list.Len()
	for ii := 0; ii < listLen; ii++ {
		p := list.At(ii)
		name := p.Name()
		if name == "" {
			name = fmt.Sprintf(nameFormat, ii+1)
		}
		typename := types.TypeString(p.Type(), m.packageQualifier)
		// check for final variadic argument
		variadic := sig.Variadic() && ii == listLen-1 && typename[0:2] == "[]"
		param := &param{
			Name:     name,
			Type:     typename,
			Variadic: variadic,
		}
		params = append(params, param)
	}
	return params
}

func pkgInfoFromPath(src string, mode packages.LoadMode) (*packages.Package, error) {
	conf := packages.Config{
		Mode: mode,
		Dir:  src,
	}
	pkgs, err := packages.Load(&conf)
	if err != nil {
		return nil, err
	}
	if len(pkgs) == 0 {
		return nil, errors.New("No packages found")
	}
	if len(pkgs) > 1 {
		return nil, errors.New("More than one package was found")
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

func (m *method) ReturnArglist() string {
	params := make([]string, len(m.Returns))
	for i, p := range m.Returns {
		params[i] = p.TypeString()
	}
	if len(m.Returns) > 1 {
		return fmt.Sprintf("(%s)", strings.Join(params, ", "))
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
