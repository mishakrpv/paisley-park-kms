package messaging

type MessageBroker interface {
	Produce() error
	Consume() error
}
