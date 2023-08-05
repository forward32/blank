package bus

import "time"

type OrderCreatedEvent struct {
	Email     string    `json:"email"`
	Product   string    `json:"product"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
