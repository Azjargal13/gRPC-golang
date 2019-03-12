package main

import (
	"context"
	"log"
	"net"

	calculator "github.com/GolangTutorials/calculator/proto"
	"google.golang.org/grpc"
)

const port = ":50051"

type server struct{}

func (*server) Sum(ctx context.Context, req *calculator.SumRequest) (*calculator.SumResponse, error) {
	fName := req.FNum
	sName := req.SNum

	sum := fName + sName
	res := &calculator.SumResponse{
		SumRes: sum,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
