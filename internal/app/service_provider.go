package app

import (
	"context"
	"log"

	"github.com/drizzleent/chat-server/internal/api/chat"
	"github.com/drizzleent/chat-server/internal/config"
	"github.com/drizzleent/chat-server/internal/config/env"
	"github.com/drizzleent/chat-server/internal/service"
	chatService "github.com/drizzleent/chat-server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig   config.PgConfig
	grpcConfig config.GRPCConfig

	chatService service.ChatService

	chatImpl *chat.Implementation
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

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if nil == s.chatService {
		s.chatService = chatService.NewService()
	}

	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.Implementation {
	if nil == s.chatImpl {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}
