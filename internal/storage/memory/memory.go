package memory

import "github.com/forward32/blank/internal/storage"

var _ storage.Storage = (*Memory)(nil)

type Memory struct{}

func (m *Memory) CreateOrder(order *storage.Order) error {
	return nil
}

func (m *Memory) UpdateStats(update *storage.StatsUpdate) error {
	return nil
}
