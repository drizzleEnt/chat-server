package chat

import "github.com/drizzleent/chat-server/internal/service"

type srv struct {
}

func NewService() service.ChatService {
	return &srv{}
}
