package chat

import (
	"github.com/drizzleent/chat-server/internal/model"
)

func (i *Implementation) ConnectChat(cl *Client) {
	i.mxChannel.RLock()
	_, ok := i.channels[cl.ChatID]
	i.mxChannel.RUnlock()

	if !ok {
		err := i.CreateChat(cl.Conn().Request().Context(), cl.ChatID)
		if err != nil {
			return
		}
		//status.Errorf(codes.NotFound, "chat not found ")
		i.mxChannel.RLock()
		i.channels[cl.ChatID] = make(chan *model.InMessage, 100)
		i.mxChannel.RUnlock()
	}

	i.mxChat.Lock()
	if _, chatOk := i.chats[cl.ChatID]; !chatOk {
		i.chats[cl.ChatID] = &Chat{
			streams: map[int]*Client{},
		}
	}
	i.mxChat.Unlock()

	i.chats[cl.ChatID].m.Lock()
	i.chats[cl.ChatID].streams[cl.ID] = cl
	i.chats[cl.ChatID].m.Unlock()
}
