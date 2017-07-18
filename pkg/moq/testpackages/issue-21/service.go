// Package moqtest is just a package to try getting started with moq without the overhead of a real package
package moqtest

//go:generate moq -out service_moq_test.go -pkg moqtest_test . Service

// Service is the interface which should be mocked by moq
type Service interface {
	User(ID string) (User, error)
}

// User is just a struct for testing
type User struct {
	Name string
}
