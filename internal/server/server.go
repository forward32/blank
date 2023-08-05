package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/forward32/blank/internal/bus"
	"github.com/forward32/blank/internal/storage"
)

type Server struct {
	s storage.Storage
	p bus.Producer
}

func New(s storage.Storage, p bus.Producer) *Server {
	return &Server{s, p}
}

func (s *Server) Run(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/order/create", s.createOrder)

	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return srv.ListenAndServe()
}

func (s *Server) createOrder(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	order := storage.Order{
		Email:     req.Email,
		EntityID:  req.EntityID,
		Price:     req.Price,
		CreatedAt: time.Now(),
	}
	if err := s.s.CreateOrder(&order); err != nil {
		http.Error(w, "failed to create order", http.StatusInternalServerError)
		return
	}

	update := storage.StatsUpdate{
		EntityID:  order.EntityID,
		Price:     order.Price,
		UpdatedAt: order.CreatedAt,
	}
	if err := s.s.UpdateStats(&update); err != nil {
		http.Error(w, "failed to update stats", http.StatusInternalServerError)
		return
	}

	event := bus.OrderCreatedEvent{
		Email:     order.Email,
		Product:   order.EntityID,
		Price:     order.Price,
		CreatedAt: order.CreatedAt,
	}
	if err := s.p.OrderCreated(&event); err != nil {
		http.Error(w, "failed to send order created event", http.StatusInternalServerError)
		return
	}

	log.Printf("order created: email=%s, entity_id=%s, price=%s",
		order.Email, order.EntityID, order.Price)
}
