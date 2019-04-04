package two

import (
	"net/http"

	myhttp "github.com/matryer/moq/pkg/moq/testpackages/imports/http"
	"github.com/matryer/moq/pkg/moq/testpackages/imports/one"
)

// DoSomething does something.
type DoSomething interface {
	Do(thing one.Thing) error
	Another(thing one.Thing) error
	Duplicate(fancy myhttp.FancyStruct, lessFancy http.Request)
}
