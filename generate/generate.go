package generate

//go:generate moq -out generated.go . MyInterface

type MyInterface interface {
	One() bool
	Two() int
	Three() string
}
