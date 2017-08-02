package example

import "context"

// Person is a person.
type Person struct {
	ID      string
	Name    string
	Company string
	Website string
}

// PersonStore stores people.
type PersonStore interface {
	Get(ctx context.Context, id string) (*Person, error)
	Create(ctx context.Context, person *Person, confirm bool) error
	ClearCache(id string)
}
