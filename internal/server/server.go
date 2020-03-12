package server

import (
	"context"
	"orihime/internal/protobuf"
	"orihime/internal/database"
)

type OrihimeServer struct {
	protobuf.UnimplementedOrihimeServer
}

func NewServer() *OrihimeServer {
	return &OrihimeServer{}
}

func (s *OrihimeServer) AddText(ctx context.Context, req *protobuf.TextToAdd) (*protobuf.TextAdded, error) {
	database.AddText(req.Content, req.Source)
	return &protobuf.TextAdded{}, nil
}

func (s *OrihimeServer) AddWord(ctx context.Context, req *protobuf.WordToAdd) (*protobuf.WordAdded, error) {
	database.AddWord(req.Word, req.DefinitionText, req.Source)
	return &protobuf.WordAdded{}, nil
}

func (s *OrihimeServer) AddSource(ctx context.Context, req *protobuf.SourceToAdd) (*protobuf.SourceAdded, error) {
	database.AddSource(req.Source)
	return &protobuf.SourceAdded{}, nil
}

func (s *OrihimeServer) AddChildWord(ctx context.Context, req *protobuf.ChildWordToAdd) (*protobuf.ChildWordAdded, error) {
	database.AddChildWord(req.Word, req.Definition, req.Source, req.User, req.ParentTextHash)
	return &protobuf.ChildWordAdded{}, nil
}
