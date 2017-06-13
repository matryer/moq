package multipackage

import "github.com/matryer/moq/package/moq/testdata/multipackage/subpackage"

type MyInterface interface {
	Method1() subpackage.Something
}