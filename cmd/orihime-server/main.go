package main

import (
	"net"
	"log"
	grpc "google.golang.org/grpc"
	"orihime/internal/protobuf"
	"orihime/internal/server"
	"orihime/internal/server/config"
)

func main() {
	log.Printf("%v", config.Config)

	lis, err := net.Listen("tcp", config.Config.Server.Listen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// grpcServer := grpc.NewServer(opts...)
	grpcServer := grpc.NewServer()
	protobuf.RegisterOrihimeServer(grpcServer, server.NewServer())
	grpcServer.Serve(lis)
}
