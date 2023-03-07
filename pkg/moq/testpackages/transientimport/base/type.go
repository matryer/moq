package base

import (
	four "github.com/matryer/moq/pkg/moq/testpackages/transientimport/four/app/v1"
	one "github.com/matryer/moq/pkg/moq/testpackages/transientimport/one/v1"
	"github.com/matryer/moq/pkg/moq/testpackages/transientimport/onev1"
	three "github.com/matryer/moq/pkg/moq/testpackages/transientimport/three/v1"
	two "github.com/matryer/moq/pkg/moq/testpackages/transientimport/two/app/v1"
)

type Transient interface {
	DoSomething(onev1.Zero, one.One, two.Two, three.Three, four.Four)
}
