package main

import (
	"log"
	"net"
	"strconv"
	"time"

	serverstream "github.com/GolangTutorials/serverStreamAPI/proto"
	"google.golang.org/grpc"
)

const port = ":50051"

type server struct{}

func (*server) Greet(req *serverstream.GreetManyRequest, stream serverstream.GreetingService_GreetServer) error {
	firstname := req.GetGreeting().Firstname
	for i := 0; i < 10; i++ {
		log.Printf("Response from stream server:" + strconv.Itoa(i))
		result := "Hello " + firstname + "number " + strconv.Itoa(i)
		res := &serverstream.GreetManyResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	log.Printf("Successfully streamed the result")
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	serverstream.RegisterGreetingServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
