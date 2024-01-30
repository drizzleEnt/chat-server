package chat

import "github.com/drizzleent/chat-server/internal/service"

type srvc struct {
}

func NewService() service.ChatService {
	return &srvc{}
}
