package service_test

import (
	"context"
	"fmt"
	"grpc-demo/proto/dog"
	"grpc-demo/repo"
	"grpc-demo/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDogService_GetDog(t *testing.T) {

	breed := "husky"

	mockResponse := &dog.DogResponse{
		ImageData: []byte("fake-image-data"),
	}

	mockApi := &repo.DogRepoMock{}

	mockApi.On("GetDog", breed).Return(mockResponse, nil)

	s := &service.DogService{
		Repo: mockApi,
	}

	result, err := s.GetDog(context.Background(), &dog.DogRequest{
		Breed: breed,
	})

	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, mockResponse, result, "Unexpected result")
}

func TestDogService_GetDog_Error(t *testing.T) {

	breed := "husky"

	var fakeError = fmt.Errorf("fake-error")

	mockApi := &repo.DogRepoMock{}

	mockApi.On("GetDog", breed).Return(&dog.DogResponse{}, fakeError)

	s := &service.DogService{
		Repo: mockApi,
	}

	_, err := s.GetDog(context.Background(), &dog.DogRequest{
		Breed: breed,
	})

	assert.Equal(t, fakeError, err, "Unexpected result")
}
