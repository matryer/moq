package withresets

import "context"

// Person is a person.
type Person struct {
	ID      string
	Name    string
	Company string
	Website string
}

//go:generate moq -with-resets -out with_resets_moq_test.go . PersonStore

// PersonStore stores people.
type PersonStore interface {
	Get(ctx context.Context, id string) (*Person, error)
	Create(ctx context.Context, person *Person, confirm bool) error
	ClearCache(id string)
}
