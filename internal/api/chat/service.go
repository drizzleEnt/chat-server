package api

import desc "github.com/drizzleent/chat-server/pkg/chat_v1"

type Implementation struct {
	desc.UnimplementedChatV1Server
}

func NewImplementation() *Implementation {
	return &Implementation{}
}
