package user

import "github.com/matryer/somerepo"

// Service does something good with computers.
type Service interface {
	DoSomething(somerepo.SomeType) error
}
