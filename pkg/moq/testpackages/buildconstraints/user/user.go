package user

import "github.com/matryer/buildconstraints"

// Service does something good with computers.
type Service interface {
	DoSomething(buildconstraints.SomeType) error
}

