package main

import (
	"fmt"
	"log"
	"net"

	"../storytel"
	"google.golang.org/grpc"
)

func main() {

	fmt.Printf("%+v\n", "running server")

	// listen on 7777 port
	grpcAddress := fmt.Sprintf("%s:%d", "localhost", 7777)

	err := startGRPCServer(grpcAddress)
	if err != nil {
		log.Fatalf("failed to start gRPC server: %s", err)
	}

}

// starts the grpc server and waits for connection
func startGRPCServer(address string) error {

	// create a listener on TCP port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// create a server instance
	s := storytel.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Courses service to the server
	storytel.RegisterCoursesServiceServer(grpcServer, &s)

	// start the server
	log.Printf("starting gRPC server on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}

	return nil
}
