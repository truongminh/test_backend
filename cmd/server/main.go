package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"grpc-demo/proto/dog"
	"grpc-demo/service"

	"google.golang.org/grpc"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	dogService := &service.DogService{}

	dog.RegisterDogServiceServer(s, dogService)

	log.Println("Starting gRPC server port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
