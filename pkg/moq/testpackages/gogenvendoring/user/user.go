package user

import "github.com/sudo-suhas/moq-test-pkgs/somerepo"

//go:generate moq -out user_moq_test.go . Service

// Service does something good with computers.
type Service interface {
	DoSomething(somerepo.SomeType) error
}
