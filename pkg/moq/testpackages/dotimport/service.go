// Package dotimport addresses issue 21.
package dotimport

//go:generate moq -out service_moq_test.go -pkg dotimport_test . Service

// Service is the interface which should be mocked by moq
type Service interface {
	User(ID string) (User, error)
}

// User is just a struct for testing
type User struct {
	Name string
}
