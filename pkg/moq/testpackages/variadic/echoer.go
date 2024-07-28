package variadic

// Echoer is an interface.
type Echoer interface {
	Echo(ss ...string) []string
}
