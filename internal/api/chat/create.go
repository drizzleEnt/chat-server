package api

import (
	"context"

	"github.com/drizzleent/chat-server/internal/converter"
	desc "github.com/drizzleent/chat-server/pkg/chat_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.chatService.Create(ctx, converter.ToChatFromDescCreate(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
