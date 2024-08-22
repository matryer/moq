package generics

import (
	"context"
	"fmt"
)

// GenericStore1 is a test interface.
type GenericStore1[T Key1, S any] interface {
	Get(ctx context.Context, id T) (S, error)
	Create(ctx context.Context, id T, value S) error
}

// GenericStore2 is a test interface.
type GenericStore2[T Key2, S any] interface {
	Get(ctx context.Context, id T) (S, error)
	Create(ctx context.Context, id T, value S) error
}

// AliasStore is a test interface.
type AliasStore GenericStore1[KeyImpl, bool]

// Key1 is a test interface.
type Key1 interface {
	fmt.Stringer
}

// Key2 is a test interface.
type Key2 interface {
	~[]byte | string
}

// KeyImpl is a test type.
type KeyImpl []byte

func (x KeyImpl) String() string {
	return string(x)
}
