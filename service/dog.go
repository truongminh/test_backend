package service

import (
	"context"
	"fmt"
	"grpc-demo/api"
	"grpc-demo/proto/dog"
)

type DogService struct {
	dog.UnimplementedDogServiceServer
	Api api.IApi
}

func (s *DogService) GetDog(ctx context.Context, req *dog.DogRequest) (*dog.DogResponse, error) {

	fmt.Println("breed:", req.Breed)

	return s.Api.GetDog(req.Breed)
}
