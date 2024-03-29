package chat

import (
	"context"

	desc "github.com/drizzleent/chat-server/pkg/chat_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (i *Implementation) CreateChat(ctx context.Context, req *empty.Empty) (*desc.CreateChatResponse, error) {

	chatId, err := i.chatService.CreateChat(ctx)
	if err != nil {
		return nil, err
	}

	i.channels[chatId.String()] = make(chan *desc.Message, 100)

	return &desc.CreateChatResponse{
		ChatId: chatId.String(),
	}, nil
}
