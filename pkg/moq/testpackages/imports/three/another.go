package two

import (
	"github.com/matryer/moq/pkg/moq/testpackages/imports/another/one"
)

// DoAnother does the other thing.
type DoAnother interface {
	Do(thing one.Thing) error
}
