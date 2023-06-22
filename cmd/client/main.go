package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"grpc-demo/proto/dog"

	"google.golang.org/grpc"
)

func main() {

	address := os.Getenv("ADDRESS")

	if address == "" {
		address = "localhost:50051"
	}

	breed := flag.String("breed", "boxer", "The breed name")
	output := flag.String("output", "breed_image.jpg", "Save the image on this local file")

	flag.Parse()

	fmt.Println("server address:", address)

	// dial connection
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// create the client
	client := dog.NewDogServiceClient(conn)

	// now, call the request image API
	req := &dog.DogRequest{
		Breed: *breed,
	}
	resp, err := client.GetDog(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get image: %v", err)
	}

	// save the image on a local file
	err = ioutil.WriteFile(*output, resp.ImageData, os.ModePerm)
	if err != nil {
		log.Fatalf("could not save image: %v", err)
	}
	log.Printf("Breed %s's image was successfully saved to %s ", *breed, *output)
}
