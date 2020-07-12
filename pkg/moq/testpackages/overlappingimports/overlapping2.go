package template

import "github.com/matryer/moq/pkg/moq/testpackages/overlappingimports/one/one"

type Example2 interface {
	Example2MultiFile() one.Example
}
