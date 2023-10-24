package api

import (
	"context"
)

type keysServer struct {
	UnimplementedKeysServer
}

func NewKeysServer() *keysServer {
	return &keysServer{}
}

func (s *keysServer) GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error) {
	return &GetValueResponse{
		Value: "unknown",
	}, nil
}
