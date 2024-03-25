package chat

import (
	"context"

	"github.com/google/uuid"
)

func (s *srv) CreateChat(ctx context.Context) (uuid.UUID, error) {
	chatId, err := uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, err
	}

	_, err = s.repo.CreateChat(ctx, chatId)
	if err != nil {
		return uuid.UUID{}, err
	}
	return chatId, nil
}
