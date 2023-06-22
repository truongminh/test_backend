package service

import (
	"context"
	"fmt"
	"grpc-demo/proto/dog"
	"grpc-demo/repo"
)

type DogService struct {
	dog.UnimplementedDogServiceServer
	Repo repo.IDogRepo
}

func (s *DogService) GetDog(ctx context.Context, req *dog.DogRequest) (*dog.DogResponse, error) {

	fmt.Println("breed:", req.Breed)

	return s.Repo.GetDog(req.Breed)
}
