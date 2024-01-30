package app

import (
	"log"

	"github.com/drizzleent/chat-server/internal/config"
	"github.com/drizzleent/chat-server/internal/config/env"
)

type serviceProvider struct {
	pgConfig   config.PgConfig
	grpcConfig config.GRPCConfig
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PgConfig {
	if nil == s.pgConfig {
		cfg, err := env.NewPgConfig()
		if err != nil {
			log.Fatalf("failed to load pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if nil == s.grpcConfig {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to load grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}
