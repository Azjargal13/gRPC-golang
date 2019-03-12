package main

import (
	"context"
	"log"
	"net"

	protobuf "github.com/GolangTutorials/unaryapi/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *protobuf.GreetRequest) (*protobuf.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &protobuf.GreetResponse{
		Result: result,
	}
	return res, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	s := grpc.NewServer()
	protobuf.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Faied to serve: %v", err)
	}
}
