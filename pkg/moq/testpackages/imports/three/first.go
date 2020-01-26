package two

import (
	"github.com/matryer/moq/pkg/moq/testpackages/imports/one"
)

// DoFirst does the first thing.
type DoFirst interface {
	Do(thing one.Thing) error
}
