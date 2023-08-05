package server

type CreateOrderRequest struct {
	Email    string `json:"email"`
	EntityID string `json:"entity_id"`
	Price    string `json:"price"`
}
