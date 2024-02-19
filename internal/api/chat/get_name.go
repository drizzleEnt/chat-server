package chat

import (
	"context"

	desc "github.com/drizzleent/chat-server/pkg/chat_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetName(ctx context.Context, empty *empty.Empty) (*desc.GetNameResponse, error) {
	res, err := i.chatService.GetName(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	return &desc.GetNameResponse{
		Name: res,
	}, nil
}
