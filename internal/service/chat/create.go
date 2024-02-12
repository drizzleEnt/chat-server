package chat

import (
	"context"

	"github.com/drizzleent/chat-server/internal/model"
	auth "github.com/drizzleent/chat-server/pkg/user_v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *srv) Create(ctx context.Context, info *model.Chat) (int64, error) {
	//id, err := s.repo.Create(ctx, info)
	cl, err := connectAuthServer()
	if err != nil {
		return 0, err
	}

	resp, err := cl.Create(ctx, &auth.CreateRequest{
		Info: &auth.UserCreate{
			UserUpdate: &auth.UserUpdate{
				Name:  info.Username,
				Email: "",
				Role:  0,
			},
			Password: info.Password,
		},
	})

	if err != nil {
		return 0, err
	}

	return resp.GetId(), nil
}

func connectAuthServer() (auth.UserV1Client, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := auth.NewUserV1Client(conn)

	return client, nil
}
