package typealiaspkgimport

import "github.com/matryer/moq/pkg/moq/testpackages/typealiaspkgimport/alias"

type Processor interface {
	Process(msg alias.Message) error
}
