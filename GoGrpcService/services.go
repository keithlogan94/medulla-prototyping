package main

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) CreateDatabase(ctx context.Context, message *CreateDatabaseRequest) (*CreateDatabaseResponse, error) {
	log.Printf("Received message body from request : %s", message.Body)
	return &CreateDatabaseResponse{Body: "Hey from Server"}, nil
}
