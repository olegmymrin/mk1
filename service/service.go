package service

import (
	"fmt"
	"log"
	"net"

	"github.com/olegmymrin/mk1/service/api"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	db     *gorm.DB
	server *grpc.Server

	keysServer api.KeysServer
}

func NewService(
	config *Config,
) *Service {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	db, err := gorm.Open(postgres.Open(connStr))
	if err != nil {
		panic(err)
	}

	log.Printf("Connected to database %s", config.DBName)

	return &Service{
		db:         db,
		server:     grpc.NewServer(),
		keysServer: api.NewKeysServer(),
	}
}

func (s *Service) ListenAndServe() {
	api.RegisterKeysServer(s.server, api.NewKeysServer())
	listener, err := net.Listen("tcp4", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	if err := s.server.Serve(listener); err != nil {
		panic(err)
	}
}
