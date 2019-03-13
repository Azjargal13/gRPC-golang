package main

import (
	"context"
	"log"

	multiplier "github.com/GolangTutorials/multiplier/proto"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := multiplier.NewMultipleServiceClient(cc)
	multPi(c)
}

func multPi(c multiplier.MultipleServiceClient) {
	req := &multiplier.MultRequest{
		Fnum: 30,
		Snum: 30,
	}
	res, err := c.Multiply(context.Background(), req)
	if err != nil {
		log.Fatalf("error while multiply: %v", err)
	}
	log.Printf("Response from Mult: %v", res.MulResult)
}
