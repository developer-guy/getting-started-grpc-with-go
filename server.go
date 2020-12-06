package main

import (
	"github.com/developer-guy/getting-started-grpc-with-go/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cs := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatsServiceServer(grpcServer, &cs)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
