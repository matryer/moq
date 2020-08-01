package channels

// Queue is a type to be sent down a channel.
type Queue []string

// Queuer provides a channel example.
type Queuer interface {
	Sub(topic string) (<-chan Queue, error)
	Unsub(topic string)
}
