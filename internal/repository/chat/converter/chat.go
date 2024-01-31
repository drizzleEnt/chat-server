package converter

import (
	"github.com/drizzleent/chat-server/internal/model"
	modelRepo "github.com/drizzleent/chat-server/internal/repository/chat/model"
)

func ToChatFromRepo(chat *modelRepo.Chat) *model.Chat {
	return &model.Chat{
		Id:   chat.Id,
		Name: chat.Name,
		Msg:  chat.Msg,
	}
}
