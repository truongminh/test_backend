package service

import (
	"context"
	"fmt"
	"grpc-demo/api"
	"grpc-demo/proto/dog"
)

type DogService struct {
	dog.UnimplementedDogServiceServer
}

func (s *DogService) GetDog(ctx context.Context, req *dog.DogRequest) (*dog.DogResponse, error) {

	fmt.Println("breed:", req.Breed)

	return api.GetDog(req.Breed)
}
