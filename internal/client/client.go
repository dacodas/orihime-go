package client

import (
	"context"
	"log"
	"time"
	"orihime/internal/protobuf"
	grpc "google.golang.org/grpc"
)

var (
	InstantiatedOrihimeClient protobuf.OrihimeClient
	serverAddress string = "localhost:12345"
	orihimeGRPCContext context.Context
	options []grpc.CallOption
)

func AddChildWord(word string, definition string, source string, user string, parentTextHash []byte) {
	InstantiatedOrihimeClient.AddChildWord(orihimeGRPCContext, &protobuf.ChildWordToAdd{Word: word, Definition: definition, Source: source, User: user, ParentTextHash: parentTextHash}, options...)
}

func AddText(contents string, source string) {
	InstantiatedOrihimeClient.AddText(orihimeGRPCContext, &protobuf.TextToAdd{Content: contents, Source: source}, options...)
}

func AddWord(word string, definitionText string, source string) {
	InstantiatedOrihimeClient.AddWord(orihimeGRPCContext, &protobuf.WordToAdd{Word: word, DefinitionText: definitionText, Source: source}, options...)
}

func AddSource(source string) {
	InstantiatedOrihimeClient.AddSource(orihimeGRPCContext, &protobuf.SourceToAdd{Source: source}, options...)
}

func init() {
	orihimeGRPCContext, _ = context.WithTimeout(context.Background(), 10*time.Second)
	options = []grpc.CallOption{}

	connection, err := grpc.Dial(serverAddress, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	// defer connection.Close()

	InstantiatedOrihimeClient = protobuf.NewOrihimeClient(connection)
}
