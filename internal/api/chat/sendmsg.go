package chat

import (
	"fmt"

	"github.com/drizzleent/chat-server/internal/model"
)

func (i *Implementation) SendMessageToClient(incomeMsg *model.InMessage) {
	i.mxChannel.RLock()
	chatChan, ok := i.channels[incomeMsg.ChatID]
	i.mxChannel.RUnlock()

	if !ok {
		fmt.Println("failed send msg: chat not found")
		return
	}

	msg := <-chatChan

	for _, chatClient := range i.chats[incomeMsg.ChatID].streams {
		req := model.OutMessage{
			From: msg.UserName,
			Text: msg.Text,
		}
		chatClient.Write(&req)
	}

}
