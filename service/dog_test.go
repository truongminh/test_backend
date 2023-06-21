package service_test

import (
	"context"
	"grpc-demo/api"
	"grpc-demo/proto/dog"
	"grpc-demo/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDogService_GetDog(t *testing.T) {

	breed := "husky"

	mockResponse := &dog.DogResponse{
		ImageData: []byte("fake-image-data"),
	}

	mockApi := &api.APIMock{}

	mockApi.On("GetDog", breed).Return(mockResponse, nil)

	s := &service.DogService{
		Api: mockApi,
	}

	result, err := s.GetDog(context.Background(), &dog.DogRequest{
		Breed: breed,
	})

	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, mockResponse, result, "Unexpected result")
}
