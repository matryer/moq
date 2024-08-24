package typealiasinternal

// we shouldn't be able to import these types directly, you need to use the alias in the parent package
type MyInternalType struct {
	Foo int
}

type MyGenericType[T any] struct {
	A T
}
