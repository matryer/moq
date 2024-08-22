package anonimport

import (
	"context"
)

// Example is a test interface.
type Example interface {
	Ctx(ctx context.Context)
}
