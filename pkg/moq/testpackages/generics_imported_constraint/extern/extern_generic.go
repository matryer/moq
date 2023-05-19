package extern

import (
	"io/fs"
	"os"

	"github.com/matryer/moq/pkg/moq/testpackages/generics_imported_constraint/extern2"
)

// validate gh package works
type Foo1 interface {
	*extern2.SomeType | *fs.PathError | os.FileMode
}

// validate with ptr works
type Foo2 interface {
	*fs.PathError | os.FileMode
}

// validate without ptr works
type Foo3 interface {
	fs.PathError | os.FileMode
}

type Local struct{}

// validate works with local extern, how deep can we go?
type Foo4 interface {
	Local | os.File
}

// validate works with extern extern tilde pointer
type Foo5 interface {
	~*extern2.SomeType | os.File
}
