package main

import (
	"context"
	"github.com/developer-guy/getting-started-grpc-with-go/chat"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatsServiceClient(conn)
	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client !"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}
