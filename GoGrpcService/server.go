package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("Failed to create Listener")
	}

	grpcServer := grpc.NewServer()

	fmt.Println("Serving golang grpc server on port 9000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve grpc server over port :9000", err)
	}

}
