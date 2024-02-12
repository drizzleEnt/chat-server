package chat

import (
	desc "github.com/drizzleent/chat-server/pkg/chat_v1"
	auth "github.com/drizzleent/chat-server/pkg/user_v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ConnectChat(req *desc.ConnectChatRequest, stream desc.ChatV1_ConnectChatServer) error {
	client, err := connectAuthServer()
	if err != nil {
		return err
	}

	client.Get(stream.Context(), &auth.GetRequest{Id: 0})

	i.mxChannel.RLock()
	chatChan, ok := i.channels[req.GetChatId()]
	i.mxChannel.RUnlock()

	if !ok {
		return status.Errorf(codes.NotFound, "chat not found")
	}

	i.mxChat.Lock()
	if _, chatOk := i.chats[req.GetChatId()]; !chatOk {
		i.chats[req.GetChatId()] = &Chat{
			streams: map[string]desc.ChatV1_ConnectChatServer{},
		}
	}
	i.mxChat.Unlock()

	i.chats[req.GetChatId()].m.Lock()
	i.chats[req.GetChatId()].streams[req.GetUsername()] = stream
	i.chats[req.GetChatId()].m.Unlock()

	for {
		select {
		case msg, okCh := <-chatChan:
			if !okCh {
				return nil
			}

			for _, st := range i.chats[req.GetChatId()].streams {
				if err := st.Send(msg); err != nil {
					return err
				}
			}
		case <-stream.Context().Done():
			i.chats[req.GetChatId()].m.Lock()
			delete(i.chats[req.GetChatId()].streams, req.GetUsername())
			i.chats[req.GetChatId()].m.Unlock()
			return nil
		}

	}
}

func connectAuthServer() (auth.UserV1Client, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := auth.NewUserV1Client(conn)

	return client, nil
}
