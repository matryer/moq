package registry

import (
	"fmt"
	"go/types"
	"regexp"
	"strings"
)

type GenericConstraint struct {
	Pkg  string
	Path string
	Name string
}

func NewGenericConstraint(constraint string) *GenericConstraint {
	return &GenericConstraint{
		Pkg:  getPkgName(constraint),
		Path: getPackagePath(constraint),
		Name: getName(constraint),
	}
}

// Underlying satisfies types.Type Underlying method
func (g GenericConstraint) Underlying() types.Type {
	return g
}

// String statisfies types.Type String method
func (g GenericConstraint) String() string {
	return g.Name
}

var appearsImportedRegex = regexp.MustCompile(`.+\/.+\.[^\.\/]`)

func ConstraintAppearsImported(constraint string) bool {
	return appearsImportedRegex.Match([]byte(constraint))
}

func getPkgName(constraint string) string {
	if i := strings.LastIndexByte(constraint, '/'); i != -1 {
		constraint = strings.TrimLeft(constraint[i:], "/")
	}

	if i := strings.LastIndexByte(constraint, '.'); i != -1 {
		constraint = constraint[:i]
	}

	return constraint
}

func getPackagePath(constraint string) string {
	if i := strings.LastIndexByte(constraint, '.'); i != -1 {
		constraint = constraint[:i]
	}

	return strings.TrimLeft(constraint, "*")
}

func getName(constraint string) string {
	var ptr bool
	if constraint[0] == '*' {
		ptr = true
	}

	if i := strings.LastIndexByte(constraint, '/'); i != -1 {
		constraint = strings.TrimPrefix(constraint[i:], "/")
	}

	if ptr {
		constraint = fmt.Sprintf("*%s", constraint)
	}

	return constraint
}
