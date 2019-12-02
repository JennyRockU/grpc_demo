package main

import (
	"fmt"
	"log"
	"net"

	"../api"
	"google.golang.org/grpc"
)

// start the grpc server and wait for connection
func main() {
	// listen on 7777 port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// create server instance
	s := api.Server{}

	// create grpc server object
	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	api.RegisterPingServer(grpcServer, &s)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}

}
