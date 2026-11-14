package methodconflict

// Service exercises generated field and helper method name collisions.
type Service interface {
	Begin() error
	BeginFunc(fn func() error) error
	Query() string
	QueryCalls() int
}
