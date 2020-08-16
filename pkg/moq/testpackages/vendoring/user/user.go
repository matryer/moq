package user

import "github.com/sudo-suhas/moq-test-pkgs/somerepo"

// Service does something good with computers.
type Service interface {
	DoSomething(somerepo.SomeType) error
}
