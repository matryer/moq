package genericsthirdparty

import (
	"context"

	"github.com/matryer/moq/pkg/moq/testpackages/genericsthirdparty/thirdparty"
)

type Item[T any] struct{}

type GenericStore interface {
	Get(ctx context.Context, item Item[thirdparty.Item])
	Create(ctx context.Context, item Item[thirdparty.Item])
}
