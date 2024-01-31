package chat

import (
	"github.com/drizzleent/chat-server/internal/repository"
	"github.com/drizzleent/chat-server/internal/service"
)

type srv struct {
	repo repository.ChatRepository
}

func NewService(repo repository.ChatRepository) service.ChatService {
	return &srv{
		repo: repo,
	}
}
