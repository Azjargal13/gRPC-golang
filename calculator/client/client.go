package main

import (
	"context"
	"io"
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
	doPrimeNumberDecomposition(c)
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
func doPrimeNumberDecomposition(c calculator.CalculatorServiceClient) {
	req := &calculator.PrimeNumberDecompositionRequest{
		Number: 33,
	}
	res, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to do PrimeNumber: %v", err)
	}
	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from PrimeNumber: %v", msg.Primefactor)
	}
}
