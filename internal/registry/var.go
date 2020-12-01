package registry

import (
	"go/types"
	"strings"
)

// Var represents a method variable/parameter.
//
// It should be created using a method scope instance.
type Var struct {
	vr         *types.Var
	imports    map[string]*Package
	moqPkgPath string

	Name string
}

// IsSlice returns whether the type (or the underlying type) is a slice.
func (v Var) IsSlice() bool {
	_, ok := v.vr.Type().Underlying().(*types.Slice)
	return ok
}

// TypeString returns the variable type with the package qualifier in the
// format 'pkg.Type'.
func (v Var) TypeString() string {
	return types.TypeString(v.vr.Type(), v.packageQualifier)
}

// packageQualifier is a types.Qualifier.
func (v Var) packageQualifier(pkg *types.Package) string {
	path := stripVendorPath(pkg.Path())
	if v.moqPkgPath != "" && v.moqPkgPath == path {
		return ""
	}

	return v.imports[path].Qualifier()
}

// generateVarName generates a name for the variable using the type
// information.
//
// Examples:
// - string -> s
// - int -> n
// - chan int -> intCh
// - []a.MyType -> myTypes
// - map[string]int -> stringToInt
// - error -> err
// - a.MyType -> myType
func generateVarName(t types.Type) string {
	if t, ok := t.(*types.Named); ok {
		name := t.Obj().Name()
		if name == "error" {
			name = "err"
		} else if name == deCapitalise(name) {
			return name + "MoqParam"
		}

		return name
	}

	nestedType := func(t types.Type) string {
		switch t := t.(type) {
		case *types.Basic:
			return t.String()

		case *types.Named:
			return t.Obj().Name()
		}
		return generateVarName(t)
	}

	switch t := t.(type) {
	case *types.Basic:
		return basicTypeVarName(t)

	case *types.Array:
		return deCapitalise(nestedType(t.Elem())) + "s"

	case *types.Slice:
		return deCapitalise(nestedType(t.Elem())) + "s"

	case *types.Struct: // anonymous struct
		return "val"

	case *types.Pointer:
		return deCapitalise(nestedType(t.Elem()))

	case *types.Signature:
		return "fn"

	case *types.Interface: // anonymous interface
		return "ifaceVal"

	case *types.Map:
		return deCapitalise(nestedType(t.Key())) + "To" + capitalise(nestedType(t.Elem()))

	case *types.Chan:
		return deCapitalise(nestedType(t.Elem())) + "Ch"
	}

	return "v"
}

func basicTypeVarName(b *types.Basic) string {
	switch b.Info() {
	case types.IsBoolean:
		return "b"

	case types.IsInteger:
		return "n"

	case types.IsFloat:
		return "f"

	case types.IsString:
		return "s"
	}

	return "v"
}

func capitalise(s string) string   { return strings.ToUpper(s[:1]) + s[1:] }
func deCapitalise(s string) string { return strings.ToLower(s[:1]) + s[1:] }