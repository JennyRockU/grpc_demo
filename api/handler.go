package api

import (
	fmt "fmt"

	"golang.org/x/net/context"
)

// an abstract represntation of Server type, for DI
type Server struct {
}

// SayHello creates a response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	fmt.Printf("Received message: %s\n", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}
