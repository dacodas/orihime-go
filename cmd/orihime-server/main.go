package main

import (
	"net"
	"log"
	grpc "google.golang.org/grpc"
	"orihime/internal/protobuf"
	"orihime/internal/server"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// grpcServer := grpc.NewServer(opts...)
	grpcServer := grpc.NewServer()
	protobuf.RegisterOrihimeServer(grpcServer, server.NewServer())
	grpcServer.Serve(lis)
}
