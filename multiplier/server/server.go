package main

import (
	"context"
	"log"
	"net"

	multiplier "github.com/GolangTutorials/multiplier/proto"
	"google.golang.org/grpc"
)

const port = ":50051"

type server struct{}

func (*server) Multiply(ctx context.Context, req *multiplier.MultRequest) (*multiplier.MultResponse, error) {
	fnum := req.Fnum
	snum := req.Snum

	mult := fnum * snum
	res := &multiplier.MultResponse{
		MulResult: mult,
	}
	log.Printf("Received numbers: %v, %v", fnum, snum)
	return res, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	multiplier.RegisterMultipleServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
