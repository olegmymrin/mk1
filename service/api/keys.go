package api

import (
	"context"
)

type server struct {
	UnimplementedKeysServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error) {
	return &GetValueResponse{
		Value: "unknown",
	}, nil
}
