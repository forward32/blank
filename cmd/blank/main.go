package main

import (
	"log"

	"github.com/forward32/blank/internal/bus/kafka"
	"github.com/forward32/blank/internal/server"
	"github.com/forward32/blank/internal/storage/memory"
)

func main() {
	storage := &memory.Memory{}
	producer := &kafka.Producer{}
	server := server.New(storage, producer)
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
