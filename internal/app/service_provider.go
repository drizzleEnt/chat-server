package app

import (
	"context"
	"log"

	"github.com/drizzleent/chat-server/internal/api/chat"
	"github.com/drizzleent/chat-server/internal/client/db"
	"github.com/drizzleent/chat-server/internal/client/db/pg"
	"github.com/drizzleent/chat-server/internal/config"
	"github.com/drizzleent/chat-server/internal/config/env"
	"github.com/drizzleent/chat-server/internal/repository"
	repoChat "github.com/drizzleent/chat-server/internal/repository/chat"
	"github.com/drizzleent/chat-server/internal/service"
	chatService "github.com/drizzleent/chat-server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig   config.PgConfig
	grpcConfig config.GRPCConfig

	dbClient db.Client

	chatRepository repository.ChatRepository

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

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if nil == s.dbClient {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %s", err.Error())
		}

		err = cl.DB().Ping(ctx)

		if err != nil {
			log.Fatalf("failed to ping db client: %s", err.Error())
		}

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if nil == s.chatRepository {
		s.chatRepository = repoChat.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if nil == s.chatService {
		s.chatService = chatService.NewService(s.ChatRepository(ctx))
	}

	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.Implementation {
	if nil == s.chatImpl {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}
