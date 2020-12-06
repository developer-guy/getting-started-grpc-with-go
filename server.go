package main

import (
	"github.com/developer-guy/getting-started-grpc-with-go/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cs := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatsServiceServer(grpcServer, &cs)
	reflection.Register(grpcServer)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Printf("gRPC server listening on port 9000")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Printf("gRPC server shutting down...")
}
