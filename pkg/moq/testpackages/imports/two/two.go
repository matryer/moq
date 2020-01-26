package two

import (
	another "github.com/matryer/moq/pkg/moq/testpackages/imports/another/one"
	"github.com/matryer/moq/pkg/moq/testpackages/imports/one"
)

// DoSomething does something.
type DoSomething interface {
	Do(thing one.Thing) error
	Another(thing another.Thing) error
}
