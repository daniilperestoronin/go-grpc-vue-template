package services

import (
	"log"

	"github.com/daniilperestoronin/go-grpc-vue-template/api"
	"golang.org/x/net/context"
)

// PingService represents the gRPC server
type PingService struct {
}

// SayHello generates response to a Ping request
func (s *PingService) SayHello(ctx context.Context, in *api.PingMessage) (*api.PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &api.PingMessage{Greeting: "Hello From Server"}, nil
}
