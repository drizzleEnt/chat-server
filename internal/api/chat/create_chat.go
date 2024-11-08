package chat

import (
	"context"
	"fmt"

	"github.com/drizzleent/chat-server/internal/model"
)

func (i *Implementation) CreateChat(ctx context.Context, chatID string) error {
	var chatId string
	isExist, err := i.chatService.GetChat(ctx, chatID)
	if err != nil {
		return err
	}

	if !isExist {
		fmt.Printf("chat not found, creating new chat\n")
		newChatId, err := i.chatService.CreateChat(ctx)
		if err != nil {
			return err
		}
		chatId = newChatId.String()
	}

	i.mxChannel.Lock()
	i.channels[chatId] = make(chan *model.InMessage, 100)
	i.mxChannel.Unlock()
	return nil
}
