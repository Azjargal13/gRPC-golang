package main

import (
	"context"
	"log"

	calculator "github.com/GolangTutorials/calculator/proto"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculator.NewCalculatorServiceClient(cc)
	uniryPi(c)
}
func uniryPi(c calculator.CalculatorServiceClient) {
	req := &calculator.SumRequest{
		FNum: 10,
		SNum: 15,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calc Sum: %v", err)
	}
	log.Printf("Response from Sum: %v", res.SumRes)
}
