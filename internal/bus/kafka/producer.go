package kafka

import "github.com/forward32/blank/internal/bus"

var _ bus.Producer = (*Producer)(nil)

type Producer struct{}

func (p *Producer) OrderCreated(event *bus.OrderCreatedEvent) error {
	return nil
}
