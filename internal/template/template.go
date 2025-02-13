package template

import (
	"io"
	"strings"
	"text/template"

	"github.com/matryer/moq/internal/registry"
)

// Template is the Moq template. It is capable of generating the Moq
// implementation for the given template.Data.
type Template struct {
	tmpl *template.Template
}

// New returns a new instance of Template.
func New() (Template, error) {
	tmpl, err := template.New("moq").Funcs(templateFuncs).Parse(moqTemplate)
	if err != nil {
		return Template{}, err
	}

	return Template{tmpl: tmpl}, nil
}

// Execute generates and writes the Moq implementation for the given
// data.
func (t Template) Execute(w io.Writer, data Data) error {
	return t.tmpl.Execute(w, data)
}

// moqTemplate is the template for mocked code.
// language=GoTemplate
var moqTemplate = `{{ with .BuildTag }}
//go:build {{ . }}

{{ end -}}
// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package {{.PkgName}}

import (
{{- range .Imports}}
	{{. | ImportStatement}}
{{- end}}
)

{{range $i, $mock := .Mocks -}}

{{- if not $.SkipEnsure -}}
// Ensure, that {{.MockName}} does implement {{$.SrcPkgQualifier}}{{.InterfaceName}}.
// If this is not the case, regenerate this file with moq.
var _ {{$.SrcPkgQualifier}}{{.InterfaceName -}}
	{{- if .TypeParams }}[
		{{- range $index, $param := .TypeParams}}
			{{- if $index}}, {{end -}}
			{{if $param.Constraint}}{{$param.Constraint.String}}{{else}}{{$param.TypeString}}{{end}}
		{{- end -}}
		]
	{{- end }} = &{{.MockName}}
	 {{- if .TypeParams }}[
		{{- range $index, $param := .TypeParams}}
			{{- if $index}}, {{end -}}
			{{if $param.Constraint}}{{$param.Constraint.String}}{{else}}{{$param.TypeString}}{{end}}
		{{- end -}}
		]
	{{- end -}}
{}
{{- end}}

// {{.MockName}} is a mock implementation of {{$.SrcPkgQualifier}}{{.InterfaceName}}.
//
//	func TestSomethingThatUses{{.InterfaceName}}(t *testing.T) {
//
//		// make and configure a mocked {{$.SrcPkgQualifier}}{{.InterfaceName}}
//		mocked{{.InterfaceName}} := &{{.MockName}}{
			{{- range .Methods}}
//			{{.Name}}Func: func({{.ArgList}}) {{.ReturnArgTypeList}} {
//				panic("mock out the {{.Name}} method")
//			},
			{{- end}}
//		}
//
//		// use mocked{{.InterfaceName}} in code that requires {{$.SrcPkgQualifier}}{{.InterfaceName}}
//		// and then make assertions.
//
//	}
type {{.MockName}}
{{- if .TypeParams -}}
	[{{- range $index, $param := .TypeParams}}
			{{- if $index}}, {{end}}{{$param.Name | Exported}} {{$param.TypeString}}
	{{- end -}}]
{{- end }} struct {
{{- range .Methods}}
	// {{.Name}}Func mocks the {{.Name}} method.
	{{.Name}}Func func({{.ArgList}}) {{.ReturnArgTypeList}}
{{end}}
	// calls tracks calls to the methods.
	calls struct {
{{- range .Methods}}
		// {{.Name}} holds details about calls to the {{.Name}} method.
		{{.Name}} []struct {
			{{- range .Params}}
			// {{.Name | Exported}} is the {{.Name}} argument value.
			{{.Name | Exported}} {{.TypeString}}
			{{- end}}
		}
{{- end}}
	}
{{- range .Methods}}
	lock{{.Name}} {{$.Imports | SyncPkgQualifier}}.RWMutex
{{- end}}
}
{{range .Methods}}
// {{.Name}} calls {{.Name}}Func.
func (mock *{{$mock.MockName}}
{{- if $mock.TypeParams -}}
	[{{- range $index, $param := $mock.TypeParams}}
		{{- if $index}}, {{end}}{{$param.Name | Exported}}
	{{- end -}}]
{{- end -}}
) {{.Name}}({{.ArgList}}) {{.ReturnArgTypeList}} {
{{- if not $.StubImpl}}
	if mock.{{.Name}}Func == nil {
		panic("{{$mock.MockName}}.{{.Name}}Func: method is nil but {{$mock.InterfaceName}}.{{.Name}} was just called")
	}
{{- end}}
	callInfo := struct {
		{{- range .Params}}
		{{.Name | Exported}} {{.TypeString}}
		{{- end}}
	}{
		{{- range .Params}}
		{{.Name | Exported}}: {{.Name}},
		{{- end}}
	}
	mock.lock{{.Name}}.Lock()
	mock.calls.{{.Name}} = append(mock.calls.{{.Name}}, callInfo)
	mock.lock{{.Name}}.Unlock()
{{- if .Returns}}
	{{- if $.StubImpl}}
	if mock.{{.Name}}Func == nil {
		var (
		{{- range .Returns}}
			{{.Name}} {{.TypeString}}
		{{- end}}
		)
		return {{.ReturnArgNameList}}
	}
	{{- end}}
	return mock.{{.Name}}Func({{.ArgCallList}})
{{- else}}
	{{- if $.StubImpl}}
	if mock.{{.Name}}Func == nil {
		return
	}
	{{- end}}
	mock.{{.Name}}Func({{.ArgCallList}})
{{- end}}
}

// {{.Name}}Calls gets all the calls that were made to {{.Name}}.
// Check the length with:
//
//	len(mocked{{$mock.InterfaceName}}.{{.Name}}Calls())
func (mock *{{$mock.MockName}}
{{- if $mock.TypeParams -}}
	[{{- range $index, $param := $mock.TypeParams}}
		{{- if $index}}, {{end}}{{$param.Name | Exported}}
	{{- end -}}]
{{- end -}}
) {{.Name}}Calls() []struct {
		{{- range .Params}}
		{{.Name | Exported}} {{.TypeString}}
		{{- end}}
	} {
	var calls []struct {
		{{- range .Params}}
		{{.Name | Exported}} {{.TypeString}}
		{{- end}}
	}
	mock.lock{{.Name}}.RLock()
	calls = mock.calls.{{.Name}}
	mock.lock{{.Name}}.RUnlock()
	return calls
}
{{- if $.WithResets}}
// Reset{{.Name}}Calls reset all the calls that were made to {{.Name}}.
func (mock *{{$mock.MockName}}
{{- if $mock.TypeParams -}}
	[{{- range $index, $param := $mock.TypeParams}}
		{{- if $index}}, {{end}}{{$param.Name | Exported}}
	{{- end -}}]
{{- end -}}
) Reset{{.Name}}Calls() {
	mock.lock{{.Name}}.Lock()
	mock.calls.{{.Name}} = nil
	mock.lock{{.Name}}.Unlock()
}
{{end}}
{{end -}}
{{- if $.WithResets}}
// ResetCalls reset all the calls that were made to all mocked methods.
func (mock *{{$mock.MockName}}
{{- if $mock.TypeParams -}}
	[{{- range $index, $param := $mock.TypeParams}}
		{{- if $index}}, {{end}}{{$param.Name | Exported}}
	{{- end -}}]
{{- end -}}
) ResetCalls() {
	{{- range .Methods}}
	mock.lock{{.Name}}.Lock()
	mock.calls.{{.Name}} = nil
	mock.lock{{.Name}}.Unlock()
	{{end -}}
}
{{end -}}
{{end -}}
`

// This list comes from the golint codebase. Golint will complain about any of
// these being mixed-case, like "Id" instead of "ID".
var golintInitialisms = []string{
	"ACL", "API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS",
	"QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SQL", "SSH", "TCP", "TLS", "TTL", "UDP", "UI", "UID", "UUID", "URI",
	"URL", "UTF8", "VM", "XML", "XMPP", "XSRF", "XSS",
}

var templateFuncs = template.FuncMap{
	"ImportStatement": func(imprt *registry.Package) string {
		if imprt.Alias == "" {
			return `"` + imprt.Path() + `"`
		}
		return imprt.Alias + ` "` + imprt.Path() + `"`
	},
	"SyncPkgQualifier": func(imports []*registry.Package) string {
		for _, imprt := range imports {
			if imprt.Path() == "sync" {
				return imprt.Qualifier()
			}
		}

		return "sync"
	},
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
