package generate

// In a terminal, run `go generate` in this directory to have
// it generates the generated.go file.

//go:generate moq -out generated.go . MyInterface

type MyInterface interface {
	One() bool
	Two() int
	Three() string
}
