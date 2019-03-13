package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientstream "github.com/GolangTutorials/clientStreamAPI/proto"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect... %v", err)
	}
	defer cc.Close()

	c := clientstream.NewLongGreetClient(cc)
	doSendReq(c)

}
func doSendReq(c clientstream.LongGreetClient) {
	log.Printf("Stream sent to server")
	reqs := []*clientstream.LongManyGreetRequest{
		&clientstream.LongManyGreetRequest{
			&clientstream.Greeting{
				Firstname: "Happy",
			},
		},
		&clientstream.LongManyGreetRequest{
			&clientstream.Greeting{
				Firstname: "Day",
			},
		},
		&clientstream.LongManyGreetRequest{
			&clientstream.Greeting{
				Firstname: "Laugh",
			},
		},
		&clientstream.LongManyGreetRequest{
			&clientstream.Greeting{
				Firstname: "Smile",
			},
		},
		&clientstream.LongManyGreetRequest{
			&clientstream.Greeting{
				Firstname: "Dream",
			},
		},
	}
	stream, err := c.LongManyGreet(context.Background())
	if err != nil {
		log.Fatalf("could not get strean... %v", err)
	}
	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
		log.Printf("Sending %v ", req)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not close stream... %v", err)
	}
	fmt.Printf("Response: %v\n", res)
}
