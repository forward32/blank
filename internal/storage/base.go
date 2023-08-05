package storage

type Storage interface {
	CreateOrder(order *Order) error
	UpdateStats(update *StatsUpdate) error
}
