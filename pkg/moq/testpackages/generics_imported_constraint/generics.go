package generics_imported_constraint

import (
	"context"

	"github.com/matryer/moq/pkg/moq/testpackages/generics_imported_constraint/extern"
)

//go:generate moq -out generics_moq_test.go -pkg generics_moq_test . GenericStore1

type GenericStore1[T extern.Foo1, J extern.Foo2, L extern.Foo3, F extern.Foo4, E extern.Foo5] interface {
	Tet(ctx context.Context, handler T) error
	Jet(ctx context.Context, handler J) error
	Let(ctx context.Context, handler L) error
	Fet(ctx context.Context, handler F) error
	Eet(ctx context.Context, handler E) error
}
