package moq

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"os"
	"strings"
	"text/template"
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
	src     string
	tmpl    *template.Template
	fset    *token.FileSet
	pkgs    map[string]*ast.Package
	pkgName string

	imports map[string]bool
}

// New makes a new Mocker for the specified package directory.
func New(src, packageName string) (*Mocker, error) {
	fset := token.NewFileSet()
	noTestFiles := func(i os.FileInfo) bool {
		return !strings.HasSuffix(i.Name(), "_test.go")
	}
	pkgs, err := parser.ParseDir(fset, src, noTestFiles, parser.SpuriousErrors)
	if err != nil {
		return nil, err
	}
	if len(packageName) == 0 {
		for pkgName := range pkgs {
			if strings.Contains(pkgName, "_test") {
				continue
			}
			packageName = pkgName
			break
		}
	}
	if len(packageName) == 0 {
		return nil, errors.New("failed to determine package name")
	}
	tmpl, err := template.New("moq").Funcs(templateFuncs).Parse(moqTemplate)
	if err != nil {
		return nil, err
	}
	return &Mocker{
		src:     src,
		tmpl:    tmpl,
		fset:    fset,
		pkgs:    pkgs,
		pkgName: packageName,
		imports: make(map[string]bool),
	}, nil
}

// Mock generates a mock for the specified interface name.
func (m *Mocker) Mock(w io.Writer, name ...string) error {
	if len(name) == 0 {
		return errors.New("must specify one interface")
	}
	doc := doc{
		PackageName: m.pkgName,
		Imports:     moqImports,
	}
	mocksMethods := false
	for _, pkg := range m.pkgs {
		i := 0
		files := make([]*ast.File, len(pkg.Files))
		for _, file := range pkg.Files {
			files[i] = file
			i++
		}
		conf := types.Config{Importer: newImporter(m.src)}
		tpkg, err := conf.Check(m.src, m.fset, files, nil)
		if err != nil {
			return err
		}
		for _, n := range name {
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
			}
			for i := 0; i < iiface.NumMethods(); i++ {
				mocksMethods = true
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
	}
	if mocksMethods {
		doc.Imports = append(doc.Imports, "sync")
	}
	for pkgToImport := range m.imports {
		doc.Imports = append(doc.Imports, pkgToImport)
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

func (m *Mocker) packageQualifier(pkg *types.Package) string {
	if m.pkgName == pkg.Name() {
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

type doc struct {
	PackageName string
	Objects     []obj
	Imports     []string
}

type obj struct {
	InterfaceName string
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
