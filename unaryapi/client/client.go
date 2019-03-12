package main

import (
	"context"
	"log"

	protobuf "github.com/GolangTutorials/unaryapi/proto"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer cc.Close()

	c := protobuf.NewGreetServiceClient(cc)
	unaryPi(c)

}
func unaryPi(c protobuf.GreetServiceClient) {
	req := &protobuf.GreetRequest{
		Greeting: &protobuf.Greeting{
			FirstName: "Happy",
			LastName:  "Day",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet: %v", err)
	}
	log.Printf("Response from greet: %v", res.Result)
}
