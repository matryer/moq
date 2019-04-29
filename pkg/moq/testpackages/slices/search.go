package slices

import "context"

// Search finds records
type Search interface {
	GetNames(ctx context.Context, search string) []string
	GetOther(ctx context.Context, search string) ([]string, error)
	GetMultipleSlices(ctx context.Context, search ...string) ([]string, []error)
	GetPointerSlices(ctx context.Context, search string) ([]*string, error)
	MultiSearch(ctx context.Context, searchTerms ...string) []string
}
