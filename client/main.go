package main

import (
	"fmt"
	"log"

	"../storytel"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error when establishing connectison: %s", err)
	}
	defer conn.Close()

	client := storytel.NewCoursesServiceClient(conn)

	response, err := client.Fetch(context.Background(), &storytel.CoursesRequest{Userid: "7622327"})
	if err != nil {
		log.Fatalf("error when calling Fetch courses: %s", err)
	}

	fmt.Printf("%+v\n", response.Courses)
}
