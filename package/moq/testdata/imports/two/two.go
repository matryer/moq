package two

import (
	"github.com/matryer/moq/package/moq/testdata/imports/one"
)

// DoSomething does something.
type DoSomething interface {
	Do(thing one.Thing) error
}
