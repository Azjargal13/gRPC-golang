package main

import (
	"fmt"
	"io"
	"log"
	"net"

	clientstream "github.com/GolangTutorials/clientStreamAPI/proto"
	"google.golang.org/grpc"
)

const port = ":50051"

type server struct{}

func (*server) LongManyGreet(stream clientstream.LongGreet_LongManyGreetServer) error {
	fmt.Printf("LongManyGreet function is invoked")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&clientstream.LongManyGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading: %v", err)
		}
		firstname := req.GetGreeting().Firstname
		result += "Hello" + firstname + "! "
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	clientstream.RegisterLongGreetServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
