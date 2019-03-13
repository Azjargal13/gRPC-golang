package main

import (
	"context"
	"io"
	"log"

	serverstream "github.com/GolangTutorials/serverStreamAPI/proto"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()
	c := serverstream.NewGreetingServiceClient(cc)
	doServerStream(c)
}

func doServerStream(c serverstream.GreetingServiceClient) {
	req := &serverstream.GreetManyRequest{
		Greeting: &serverstream.Greeting{
			Firstname: "Happy",
			Lastname:  "Day",
		},
	}
	resStream, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling rpc: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from Greet: %v", msg.Result)
	}
}
