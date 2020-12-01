package template

import (
	"go/types"
	"testing"

	"github.com/matryer/moq/internal/registry"
)

func TestTemplateFuncs(t *testing.T) {
	fn1 := templateFuncs["Exported"].(func(string) string)
	if fn1("var") != "Var" {
		t.Errorf("exported didn't work: %s", fn1("var"))
	}

	fn2 := templateFuncs["ImportStatement"].(func(*registry.Package) string)
	pkg := registry.NewPackage(types.NewPackage("xyz", "xyz"))
	if fn2(pkg) != `"xyz"` {
		t.Errorf("ImportStatement didn't work: %s", fn2(pkg))
	}

	pkg.Alias = "x"
	if fn2(pkg) != `x "xyz"` {
		t.Errorf("ImportStatement didn't work: %s", fn2(pkg))
	}
}
