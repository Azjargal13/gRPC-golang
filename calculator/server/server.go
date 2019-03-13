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
func (*server) PrimeNumberDecomposition(req *calculator.PrimeNumberDecompositionRequest, stream calculator.CalculatorService_PrimeNumberDecompositionServer) error {
	num := int64(req.Number)
	divisor := int64(2)

	for num > 1 {
		if num%divisor == 0 {
			stream.Send(&calculator.PrimeNumberDecompositionResponse{
				Primefactor: divisor,
			})
			num = num / divisor
		} else {
			divisor++
			log.Printf("Divisor has increaed to: %v", divisor)
		}
	}
	return nil
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
