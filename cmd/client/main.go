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

	breed := flag.String("breed", "boxer", "enter breed")
	output := flag.String("output", "received_image.jpg", "output")

	flag.Parse()

	fmt.Println("address:", address)

	// Kết nối đến server gRPC
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Tạo một client mới từ kết nối
	client := dog.NewDogServiceClient(conn)

	// Gửi yêu cầu lấy ảnh
	req := &dog.DogRequest{
		Breed: *breed,
	}
	resp, err := client.GetDog(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get image: %v", err)
	}

	// return

	// Lưu trữ ảnh nhận được từ phản hồi
	err = ioutil.WriteFile(*output, resp.ImageData, os.ModePerm)
	if err != nil {
		log.Fatalf("could not save image: %v", err)
	}
	log.Println("Image saved successfully")
}
