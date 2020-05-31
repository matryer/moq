package variadic

type Echoer interface {
	Echo(ss ...string) []string
}
