package generics

import (
	"context"
	"fmt"
)

type GenericStore1[T Key1, S any] interface {
	Get(ctx context.Context, id T) (S, error)
	Create(ctx context.Context, id T, value S) error
}

type GenericStore2[T Key2, S any] interface {
	Get(ctx context.Context, id T) (S, error)
	Create(ctx context.Context, id T, value S) error
}

type AliasStore GenericStore1[KeyImpl, bool]

type Key1 interface {
	fmt.Stringer
}

type Key2 interface {
	~[]byte | string
}

type KeyImpl []byte

func (x KeyImpl) String() string {
	return string(x)
}
