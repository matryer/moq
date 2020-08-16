// Package somerepo is a vendored package to test how moq deals with
// packages in the vendor package.
package somerepo

// SomeType is just some old type.
type SomeType struct {
	// Truth indicates whether true is true or not. Computers.
	Truth bool
}

type SomeService interface {
	Get() SomeType
}
