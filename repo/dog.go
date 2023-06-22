package repo

import (
	"encoding/json"
	"fmt"
	"grpc-demo/proto/dog"
	"io/ioutil"
	"net/http"
)

const (
	URL = "https://dog.ceo/api/breed"
)

type IDogRepo interface {
	GetDog(breed string) (*dog.DogResponse, error)
}

type DogRepo struct {
}

func (a *DogRepo) GetDog(breed string) (*dog.DogResponse, error) {

	client := &http.Client{}

	r, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/images/random", URL, breed), nil)
	if err != nil {
		fmt.Println("create request error:", err)
		return nil, err
	}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println("request error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("decoder json error:", err)
		return nil, err
	}

	if result.Status != "success" {
		return nil, fmt.Errorf("request dog breed error: %s", result.Message)
	}

	resp, err = http.Get(result.Message)
	if err != nil {
		fmt.Println("request image error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error download image:", resp.StatusCode)
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error read image:", err)
		return nil, err
	}

	return &dog.DogResponse{
		ImageData: data,
	}, nil
}
