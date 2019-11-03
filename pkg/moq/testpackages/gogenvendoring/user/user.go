package user

import "github.com/matryer/somerepo"

//go:generate moqit -out user_moq_test.go . Service

// Service does something good with computers.
type Service interface {
	DoSomething(somerepo.SomeType) error
}
