package storage

import "time"

type Order struct {
	Email     string
	EntityID  string
	Price     string
	CreatedAt time.Time
}

type StatsUpdate struct {
	EntityID  string
	Price     string
	UpdatedAt time.Time
}
