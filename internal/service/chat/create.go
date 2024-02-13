package chat

import (
	"context"

	"github.com/drizzleent/chat-server/internal/connect"
	"github.com/drizzleent/chat-server/internal/model"
	auth "github.com/drizzleent/chat-server/pkg/user_v2"
)

func (s *srv) Create(ctx context.Context, info *model.Chat) (int64, error) {
	//id, err := s.repo.Create(ctx, info)
	conn, err := connect.AuthServer()
	if err != nil {
		return 0, err
	}
	client := auth.NewUserV1Client(conn)
	defer conn.Close()
	resp, err := client.Create(ctx, &auth.CreateRequest{
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
