package channels

type Queue []string

type Queuer interface {
	Sub(topic string) (<-chan Queue, error)
}
