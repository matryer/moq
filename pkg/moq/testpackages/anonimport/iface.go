package anonimport

import (
	"context"
)

type Example interface {
	Ctx(ctx context.Context)
}
