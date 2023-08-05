package bus

type Producer interface {
	OrderCreated(event *OrderCreatedEvent) error
}
