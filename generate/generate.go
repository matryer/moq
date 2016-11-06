package generate

// In a terminal, run `go generate` in this directory to have
// it generates the generated.go file.

//go:generate moq -out generated.go . MyInterface

// MyInterface is a test interface.
type MyInterface interface {
	One() bool
	Two() int
	Three() string
}
