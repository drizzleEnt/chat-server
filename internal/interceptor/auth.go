package interceptor

import (
	"context"

	"github.com/drizzleent/chat-server/internal/auth"
	"github.com/drizzleent/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	if _, ok := req.(*chat_v1.CreateRequest); ok {
		return handler(ctx, req)
	}

	err := auth.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := handler(ctx, req)
	return res, err
}
