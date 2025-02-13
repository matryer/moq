package embedtypes

// Interface1 is a test interface.
// Its implementation must embed Embedded1.
type Interface1 interface {
	ExportedMethod()

	unexportedMethod1()
}

// Embedded1 has unexportedMethod1.
type Embedded1 struct{}

func (e Embedded1) unexportedMethod1() {}

// Interface2 is a test interface.
// Its implementation must embed Embedded2.
type Interface2 interface {
	ExportedMethod()

	unexportedMethod2()
}

// Embedded2 has unexportedMethod2.
type Embedded2 struct{}

func (e *Embedded2) unexportedMethod2() {}

// Interface3 is a test interface.
// Its implementation must embed Embedded1 and Embedded2.
type Interface3 interface {
	ExportedMethod()

	unexportedMethod1()
	unexportedMethod2()
}
