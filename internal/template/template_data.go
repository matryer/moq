package template

import (
	"fmt"
	"go/types"
	"strconv"
	"strings"

	"github.com/matryer/moq/internal/registry"
)

// Data is the template data used to render the Moq template.
type Data struct {
	PkgName         string
	SrcPkgQualifier string
	Imports         []*registry.Package
	Mocks           []MockData
	StubImpl        bool
	SkipEnsure      bool
	WithResets      bool
}

func (d Data) resolved() Data {
	if len(d.Mocks) == 0 {
		return d
	}
	mocks := make([]MockData, len(d.Mocks))
	for i, m := range d.Mocks {
		mocks[i] = m.resolved(d.WithResets)
	}
	d.Mocks = mocks
	return d
}

// MocksSomeMethod returns true of any one of the Mocks has at least 1
// method.
func (d Data) MocksSomeMethod() bool {
	for _, m := range d.Mocks {
		if len(m.Methods) > 0 {
			return true
		}
	}

	return false
}

// MockData is the data used to generate a mock for some interface.
type MockData struct {
	InterfaceName        string
	MockName             string
	TypeParams           []TypeParamData
	Methods              []MethodData
	callsFieldName       string
	resetCallsMethodName string
}

// CallsFieldName returns the struct field name used to track method calls.
func (m MockData) CallsFieldName() string {
	if m.callsFieldName != "" {
		return m.callsFieldName
	}
	return "calls"
}

// ResetCallsMethodName returns the method name used to reset all recorded calls.
func (m MockData) ResetCallsMethodName() string {
	if m.resetCallsMethodName != "" {
		return m.resetCallsMethodName
	}
	return "ResetCalls"
}

func (m MockData) resolved(withResets bool) MockData {
	methods := make([]MethodData, len(m.Methods))
	copy(methods, m.Methods)
	m.Methods = methods

	taken := make(map[string]bool, len(methods)*4+2)
	for _, method := range methods {
		taken[method.Name] = true
	}

	assignMethodNames(
		m.Methods,
		taken,
		func(name string) (string, string) { return name + "Func", name + "MoqFunc" },
		func(md *MethodData, resolved string) { md.funcFieldName = resolved },
	)

	assignMethodNames(
		m.Methods,
		taken,
		func(name string) (string, string) { return name + "Calls", name + "MoqCalls" },
		func(md *MethodData, resolved string) { md.callsMethodName = resolved },
	)

	if withResets {
		assignMethodNames(
			m.Methods,
			taken,
			func(name string) (string, string) { return "Reset" + name + "Calls", "Reset" + name + "MoqCalls" },
			func(md *MethodData, resolved string) { md.resetCallsMethodName = resolved },
		)
		m.resetCallsMethodName = uniqueIdentifier(taken, "ResetCalls", "ResetMoqCalls")
	}

	m.callsFieldName = uniqueIdentifier(taken, "calls", "moqCalls")

	assignMethodNames(
		m.Methods,
		taken,
		func(name string) (string, string) { return "lock" + name, "moqLock" + name },
		func(md *MethodData, resolved string) { md.lockFieldName = resolved },
	)

	return m
}

func assignMethodNames(
	methods []MethodData,
	taken map[string]bool,
	candidates func(name string) (primary, fallback string),
	set func(md *MethodData, resolved string),
) {
	collided := make([]bool, len(methods))
	for i := range methods {
		primary, _ := candidates(methods[i].Name)
		if !taken[primary] {
			taken[primary] = true
			set(&methods[i], primary)
		} else {
			collided[i] = true
		}
	}
	for i := range methods {
		if !collided[i] {
			continue
		}
		primary, fallback := candidates(methods[i].Name)
		set(&methods[i], uniqueIdentifier(taken, primary, fallback))
	}
}

func uniqueIdentifier(taken map[string]bool, primary, fallback string) string {
	if !taken[primary] {
		taken[primary] = true
		return primary
	}
	if !taken[fallback] {
		taken[fallback] = true
		return fallback
	}
	for n := 1; ; n++ {
		candidate := fallback + strconv.Itoa(n)
		if !taken[candidate] {
			taken[candidate] = true
			return candidate
		}
	}
}

// MethodData is the data which represents a method on some interface.
type MethodData struct {
	Name                 string
	Params               []ParamData
	Returns              []ParamData
	funcFieldName        string
	callsMethodName      string
	resetCallsMethodName string
	lockFieldName        string
}

// FuncFieldName returns the struct field name used to mock the method.
func (m MethodData) FuncFieldName() string {
	if m.funcFieldName != "" {
		return m.funcFieldName
	}
	return m.Name + "Func"
}

// CallsMethodName returns the method name used to get recorded calls for the method.
func (m MethodData) CallsMethodName() string {
	if m.callsMethodName != "" {
		return m.callsMethodName
	}
	return m.Name + "Calls"
}

// ResetCallsMethodName returns the method name used to reset recorded calls for the method.
func (m MethodData) ResetCallsMethodName() string {
	if m.resetCallsMethodName != "" {
		return m.resetCallsMethodName
	}
	return "Reset" + m.Name + "Calls"
}

// LockFieldName returns the mutex field name used to synchronize calls to the method.
func (m MethodData) LockFieldName() string {
	if m.lockFieldName != "" {
		return m.lockFieldName
	}
	return "lock" + m.Name
}

// ArgList is the string representation of method parameters, ex:
// 's string, n int, foo bar.Baz'.
func (m MethodData) ArgList() string {
	params := make([]string, len(m.Params))
	for i, p := range m.Params {
		params[i] = p.MethodArg()
	}
	return strings.Join(params, ", ")
}

// ArgCallList is the string representation of method call parameters,
// ex: 's, n, foo'. In case of a last variadic parameter, it will be of
// the format 's, n, foos...'
func (m MethodData) ArgCallList() string {
	params := make([]string, len(m.Params))
	for i, p := range m.Params {
		params[i] = p.CallName()
	}
	return strings.Join(params, ", ")
}

// ReturnArgTypeList is the string representation of method return
// types, ex: 'bar.Baz', '(string, error)'.
func (m MethodData) ReturnArgTypeList() string {
	params := make([]string, len(m.Returns))
	for i, p := range m.Returns {
		params[i] = p.TypeString()
	}
	if len(m.Returns) > 1 {
		return fmt.Sprintf("(%s)", strings.Join(params, ", "))
	}
	return strings.Join(params, ", ")
}

// ReturnArgNameList is the string representation of values being
// returned from the method, ex: 'foo', 's, err'.
func (m MethodData) ReturnArgNameList() string {
	params := make([]string, len(m.Returns))
	for i, p := range m.Returns {
		params[i] = p.Name()
	}
	return strings.Join(params, ", ")
}

// TypeParamData extends ParamData with a constraint.
type TypeParamData struct {
	ParamData
	Constraint types.Type
}

// ParamData is the data which represents a parameter to some method of
// an interface.
type ParamData struct {
	Var      *registry.Var
	Variadic bool
}

// Name returns the name of the parameter.
func (p ParamData) Name() string {
	return p.Var.Name
}

// MethodArg is the representation of the parameter in the function
// signature, ex: 'name a.Type'.
func (p ParamData) MethodArg() string {
	if p.Variadic {
		return fmt.Sprintf("%s ...%s", p.Name(), p.TypeString()[2:])
	}
	return fmt.Sprintf("%s %s", p.Name(), p.TypeString())
}

// CallName returns the string representation of the parameter to be
// used for a method call. For a variadic paramter, it will be of the
// format 'foos...'.
func (p ParamData) CallName() string {
	if p.Variadic {
		return p.Name() + "..."
	}
	return p.Name()
}

// TypeString returns the string representation of the type of the
// parameter.
func (p ParamData) TypeString() string {
	return p.Var.TypeString()
}
