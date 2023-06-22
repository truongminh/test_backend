package repo

import (
	"grpc-demo/proto/dog"

	"github.com/stretchr/testify/mock"
)

type DogRepoMock struct {
	mock.Mock
}

func (a *DogRepoMock) GetDog(breed string) (*dog.DogResponse, error) {
	args := a.Called(breed)
	return args.Get(0).(*dog.DogResponse), args.Error(1)

}
