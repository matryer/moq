package genericconstraints

import "fmt"

// Sorter exercises the comparable constraint reported in issue #200.
type Sorter[T comparable] interface {
	Sort([]T) []T
}

// StringKey combines a type set with a method requirement.
type StringKey interface {
	~string
	fmt.Stringer
}

// Store exercises dependent and method-bearing constraints without choosing instances.
type Store[K StringKey, V ~map[K]string] interface {
	Get(K) V
}
