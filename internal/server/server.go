package server

import (
	"context"
)

type orihimeServer struct {
	UnimplementedOrihimeServer
}

func (s *orihimeServer) AddText(ctx context.Context, req *TextToAdd) (*TextAdded, error) {
	return &TextAdded{}, nil
}

func (s *orihimeServer) AddWord(ctx context.Context, req *WordToAdd) (*WordAdded, error) {
	return &WordAdded{}, nil
}

func (s *orihimeServer) AddSource(ctx context.Context, req *Source) (*SourceAdded, error) {
	return &SourceAdded{}, nil

}

func (s *orihimeServer) AddChildWord(ctx context.Context, req *ChildWord) (*ChildWordAdded, error) {
	return &ChildWordAdded{}, nil

}
