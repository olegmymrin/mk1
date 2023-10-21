package main

import (
	"net"

	"github.com/olegmymrin/mk1/service/api"
	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	api.RegisterKeysServer(srv, api.NewServer())
	listener, err := net.Listen("tcp4", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
