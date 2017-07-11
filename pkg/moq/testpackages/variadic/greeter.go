package variadic

import "context"

// Greeter greets people.
type Greeter interface {
	Greet(ctx context.Context, names ...string) string
}
