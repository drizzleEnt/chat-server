package chat

import (
	"context"

	"github.com/drizzleent/chat-server/internal/converter"
	desc "github.com/drizzleent/chat-server/pkg/chat_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendRequest) (*empty.Empty, error) {
	err := i.chatService.SendMsg(ctx, converter.ToChatFromDescSend(req))

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
