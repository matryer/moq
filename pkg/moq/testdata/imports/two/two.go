package two

import (
	"github.com/matryer/moq/pkg/moq/testdata/imports/one"
)

// DoSomething does something.
type DoSomething interface {
	Do(thing one.Thing) error
	Another(thing one.Thing) error
}
