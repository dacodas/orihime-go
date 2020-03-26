package main

import (
	"net"
	"log"
	grpc "google.golang.org/grpc"
	credentials "google.golang.org/grpc/credentials"
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

	creds, err := credentials.NewServerTLSFromFile(config.Config.Server.Certificate, config.Config.Server.Key)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	protobuf.RegisterOrihimeServer(grpcServer, server.NewServer())
	grpcServer.Serve(lis)
}
