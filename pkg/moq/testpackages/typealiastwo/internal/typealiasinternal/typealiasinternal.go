package typealiasinternal

// we shouldn't be able to import this type directly, you need to use the alias in the parent package
type MyInternalType struct {
	Foo int
}
