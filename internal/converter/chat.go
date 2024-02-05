package converter

import (
	"github.com/drizzleent/chat-server/internal/model"
	desc "github.com/drizzleent/chat-server/pkg/chat_v1"
)

func ToChatFromDescCreate(chat *desc.CreateRequest) *model.Chat {
	return &model.Chat{
		Id:   0,
		Name: chat.Usernames,
		Msg:  chat.Msg,
	}
}

func ToChatFromDescDelete(chat *desc.DeleteRequest) int64 {
	return chat.Id
}

func ToChatFromDescSend(chat *desc.SendMessageRequest) *model.Chat {
	return &model.Chat{
		Id: 0,
	}
}
