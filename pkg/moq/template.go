package moq

// moqImports are the imports all moq files get.
var moqImports = []string{}

// moqTemplate is the template for mocked code.
var moqTemplate = `// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package {{.PackageName}}
{{- $sourcePackagePrefix := .SourcePackagePrefix}}

import (
{{- range .Imports }}
	"{{.}}"
{{- end }}
)

{{ range $i, $obj := .Objects -}}
var (
{{- range .Methods }}
	lock{{$obj.MockName}}{{.Name}}	sync.RWMutex
{{- end }}
)

// Ensure, that {{.MockName}} does implement {{$sourcePackagePrefix}}{{.InterfaceName}}.
// If this is not the case, regenerate this file with moq.
var _ {{$sourcePackagePrefix}}{{.InterfaceName}} = &{{.MockName}}{}

// {{.MockName}} is a mock implementation of {{$sourcePackagePrefix}}{{.InterfaceName}}.
//
//     func TestSomethingThatUses{{.InterfaceName}}(t *testing.T) {
//
//         // make and configure a mocked {{$sourcePackagePrefix}}{{.InterfaceName}}
//         mocked{{.InterfaceName}} := &{{.MockName}}{ {{ range .Methods }}
//             {{.Name}}Func: func({{ .Arglist }}) {{.ReturnArglist}} {
// 	               panic("mock out the {{.Name}} method")
//             },{{- end }}
//         }
//
//         // use mocked{{.InterfaceName}} in code that requires {{$sourcePackagePrefix}}{{.InterfaceName}}
//         // and then make assertions.
//
//     }
type {{.MockName}} struct {
{{- range .Methods }}
	// {{.Name}}Func mocks the {{.Name}} method.
	{{.Name}}Func func({{ .Arglist }}) {{.ReturnArglist}}
{{ end }}
	// calls tracks calls to the methods.
	calls struct {
{{- range .Methods }}
		// {{ .Name }} holds details about calls to the {{.Name}} method.
		{{ .Name }} []struct {
			{{- range .Params }}
			// {{ .Name | Exported }} is the {{ .Name }} argument value.
			{{ .Name | Exported }} {{ .Type }}
			{{- end }}
		}
{{- end }}
	}
}
{{ range .Methods }}
// {{.Name}} calls {{.Name}}Func.
func (mock *{{$obj.MockName}}) {{.Name}}({{.Arglist}}) {{.ReturnArglist}} {
	if mock.{{.Name}}Func == nil {
		panic("{{$obj.MockName}}.{{.Name}}Func: method is nil but {{$obj.InterfaceName}}.{{.Name}} was just called")
	}
	callInfo := struct {
		{{- range .Params }}
		{{ .Name | Exported }} {{ .Type }}
		{{- end }}
	}{
		{{- range .Params }}
		{{ .Name | Exported }}: {{ .LocalName }},
		{{- end }}
	}
	lock{{$obj.MockName}}{{.Name}}.Lock()
	mock.calls.{{.Name}} = append(mock.calls.{{.Name}}, callInfo)
	lock{{$obj.MockName}}{{.Name}}.Unlock()
{{- if .ReturnArglist }}
	return mock.{{.Name}}Func({{.ArgCallList}})
{{- else }}
	mock.{{.Name}}Func({{.ArgCallList}})
{{- end }}
}

// {{.Name}}Calls gets all the calls that were made to {{.Name}}.
// Check the length with:
//     len(mocked{{$obj.InterfaceName}}.{{.Name}}Calls())
func (mock *{{$obj.MockName}}) {{.Name}}Calls() []struct {
		{{- range .Params }}
		{{ .Name | Exported }} {{ .Type }}
		{{- end }}
	} {
	var calls []struct {
		{{- range .Params }}
		{{ .Name | Exported }} {{ .Type }}
		{{- end }}
	}
	lock{{$obj.MockName}}{{.Name}}.RLock()
	calls = mock.calls.{{.Name}}
	lock{{$obj.MockName}}{{.Name}}.RUnlock()
	return calls
}
{{ end -}}
{{ end -}}`
