package template

import (
	"go/types"
	"testing"

	"github.com/matryer/moq/internal/registry"
)

func TestTemplateFuncs(t *testing.T) {
	t.Run("Exported", func(t *testing.T) {
		f := templateFuncs["Exported"].(func(string) string)
		if f("") != "" {
			t.Errorf("Exported(...) want: ``; got: `%s`", f(""))
		}
		if f("var") != "Var" {
			t.Errorf("Exported(...) want: `Var`; got: `%s`", f("var"))
		}
	})

	t.Run("ImportStatement", func(t *testing.T) {
		f := templateFuncs["ImportStatement"].(func(*registry.Package) string)
		pkg := registry.NewPackage(types.NewPackage("xyz", "xyz"))
		if f(pkg) != `"xyz"` {
			t.Errorf("ImportStatement(...): want: `\"xyz\"`; got: `%s`", f(pkg))
		}

		pkg.Alias = "x"
		if f(pkg) != `x "xyz"` {
			t.Errorf("ImportStatement(...): want: `x \"xyz\"`; got: `%s`", f(pkg))
		}
	})

	t.Run("SyncPkgQualifier", func(t *testing.T) {
		f := templateFuncs["SyncPkgQualifier"].(func([]*registry.Package) string)
		if f(nil) != "sync" {
			t.Errorf("SyncPkgQualifier(...): want: `sync`; got: `%s`", f(nil))
		}
		imports := []*registry.Package{
			registry.NewPackage(types.NewPackage("sync", "sync")),
			registry.NewPackage(types.NewPackage("github.com/some/module", "module")),
		}
		if f(imports) != "sync" {
			t.Errorf("SyncPkgQualifier(...): want: `sync`; got: `%s`", f(imports))
		}

		syncPkg := registry.NewPackage(types.NewPackage("sync", "sync"))
		syncPkg.Alias = "stdsync"
		otherSyncPkg := registry.NewPackage(types.NewPackage("github.com/someother/sync", "sync"))
		imports = []*registry.Package{otherSyncPkg, syncPkg}
		if f(imports) != "stdsync" {
			t.Errorf("SyncPkgQualifier(...): want: `stdsync`; got: `%s`", f(imports))
		}
	})
}

func TestResolvedIdentifierConflicts(t *testing.T) {
	t.Run("BeginAndBeginFunc", func(t *testing.T) {
		d := Data{
			PkgName: "sample",
			Mocks: []MockData{
				{
					InterfaceName: "SomeIface",
					MockName:      "SomeIfaceMock",
					Methods: []MethodData{
						{Name: "Begin"},
						{Name: "BeginFunc"},
					},
				},
			},
		}.resolved()

		got := d.Mocks[0]
		if got.Methods[0].FuncFieldName() != "BeginMoqFunc" {
			t.Fatalf("Begin FuncFieldName: want BeginMoqFunc, got %s", got.Methods[0].FuncFieldName())
		}
		if got.Methods[1].FuncFieldName() != "BeginFuncFunc" {
			t.Fatalf("BeginFunc FuncFieldName: want BeginFuncFunc, got %s", got.Methods[1].FuncFieldName())
		}
		if got.CallsFieldName() != "calls" {
			t.Fatalf("CallsFieldName: want calls, got %s", got.CallsFieldName())
		}
	})

	t.Run("PreservesNonCollidingPrimaryAndResolvesChains", func(t *testing.T) {
		d := Data{
			PkgName:    "sample",
			WithResets: true,
			Mocks: []MockData{
				{
					InterfaceName: "Complex",
					MockName:      "ComplexMock",
					Methods: []MethodData{
						{Name: "Begin"},
						{Name: "BeginFunc"},
						{Name: "BeginFuncFunc"},
						{Name: "BeginMoq"},
						{Name: "Foo"},
						{Name: "FooCalls"},
						{Name: "Reset"},
						{Name: "ResetFooCalls"},
						{Name: "calls"},
						{Name: "lockFoo"},
					},
				},
			},
		}.resolved()

		mock := d.Mocks[0]
		byName := map[string]MethodData{}
		for _, m := range mock.Methods {
			byName[m.Name] = m
		}

		// BeginMoq's natural primary BeginMoqFunc does not collide with any interface method,
		// so BeginMoq keeps BeginMoqFunc while Begin steps to BeginMoqFunc1.
		if byName["BeginMoq"].FuncFieldName() != "BeginMoqFunc" {
			t.Errorf("BeginMoq FuncFieldName: want BeginMoqFunc, got %s", byName["BeginMoq"].FuncFieldName())
		}
		if byName["Begin"].FuncFieldName() != "BeginMoqFunc1" {
			t.Errorf("Begin FuncFieldName: want BeginMoqFunc1, got %s", byName["Begin"].FuncFieldName())
		}
		if byName["BeginFunc"].FuncFieldName() != "BeginFuncMoqFunc" {
			t.Errorf("BeginFunc FuncFieldName: want BeginFuncMoqFunc, got %s", byName["BeginFunc"].FuncFieldName())
		}
		if byName["BeginFuncFunc"].FuncFieldName() != "BeginFuncFuncFunc" {
			t.Errorf("BeginFuncFunc FuncFieldName: want BeginFuncFuncFunc, got %s", byName["BeginFuncFunc"].FuncFieldName())
		}
		if byName["Foo"].CallsMethodName() != "FooMoqCalls" {
			t.Errorf("Foo CallsMethodName: want FooMoqCalls, got %s", byName["Foo"].CallsMethodName())
		}
		if byName["Reset"].CallsMethodName() != "ResetCalls" {
			t.Errorf("Reset CallsMethodName: want ResetCalls, got %s", byName["Reset"].CallsMethodName())
		}
		if mock.ResetCallsMethodName() != "ResetMoqCalls" {
			t.Errorf("mock ResetCallsMethodName: want ResetMoqCalls, got %s", mock.ResetCallsMethodName())
		}
		if byName["Foo"].ResetCallsMethodName() != "ResetFooMoqCalls" {
			t.Errorf("Foo ResetCallsMethodName: want ResetFooMoqCalls, got %s", byName["Foo"].ResetCallsMethodName())
		}
		if mock.CallsFieldName() != "moqCalls" {
			t.Errorf("mock CallsFieldName: want moqCalls, got %s", mock.CallsFieldName())
		}
		if byName["Foo"].LockFieldName() != "moqLockFoo" {
			t.Errorf("Foo LockFieldName: want moqLockFoo, got %s", byName["Foo"].LockFieldName())
		}

		seen := map[string]string{}
		record := func(kind, ident string) {
			if prev, exists := seen[ident]; exists {
				t.Errorf("duplicate identifier %q used by both %s and %s", ident, prev, kind)
			}
			seen[ident] = kind
		}
		record("mock.CallsFieldName", mock.CallsFieldName())
		record("mock.ResetCallsMethodName", mock.ResetCallsMethodName())
		for _, m := range mock.Methods {
			record("method:"+m.Name, m.Name)
			record("funcField:"+m.Name, m.FuncFieldName())
			record("callsMethod:"+m.Name, m.CallsMethodName())
			record("resetMethod:"+m.Name, m.ResetCallsMethodName())
			record("lockField:"+m.Name, m.LockFieldName())
		}
	})
}
