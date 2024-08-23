package typealias

import (
	"github.com/matryer/moq/pkg/moq/testpackages/typealiastwo"
)

type Example interface {
	Do(a typealiastwo.AliasType) error
}
