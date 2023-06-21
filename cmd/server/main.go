package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"grpc-demo/api"
	"grpc-demo/proto/dog"

	"google.golang.org/grpc"
)

type server struct {
	dog.UnimplementedDogServiceServer
}

func (s *server) GetDog(ctx context.Context, req *dog.DogRequest) (*dog.DogResponse, error) {

	fmt.Println("breed:", req.Breed)

	return api.GetDog(req.Breed)
}

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
	dog.RegisterDogServiceServer(s, &server{})

	log.Println("Starting gRPC server port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
