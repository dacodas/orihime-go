package client

import (
	"log"
)

var (
	InstantiatedOrihimeClient OrihimeClient
	serverAddress string = "localhost:12345"
)

func init() {
	connection, err := grpc.Dial(serverAddress, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer connection.Close()

	InstantiatedOrihimeClient = NewOrihimeClient(connection)
}
