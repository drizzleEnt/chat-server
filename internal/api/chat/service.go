package chat

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/drizzleent/chat-server/internal/model"
	"github.com/drizzleent/chat-server/internal/service"
	"golang.org/x/net/websocket"
)

type Chat struct {
	streams map[int]*Client
	m       sync.RWMutex
}

type Implementation struct {
	chatService service.ChatService

	chats  map[string]*Chat
	mxChat sync.RWMutex

	channels  map[string]chan *model.InMessage
	mxChannel sync.RWMutex

	addCh              chan *Client
	delCh              chan *Client
	getMsgFromClientCh chan *model.InMessage
	doneCh             chan bool
	errCh              chan error
}

func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
		chats:       make(map[string]*Chat),
		channels:    make(map[string]chan *model.InMessage),

		addCh:              make(chan *Client),
		delCh:              make(chan *Client),
		getMsgFromClientCh: make(chan *model.InMessage, 100),
		doneCh:             make(chan bool),
		errCh:              make(chan error),
	}
}

func (i *Implementation) Listen() {
	onConnected := func(wsc *websocket.Conn) {
		defer func() {
			err := wsc.Close()
			if err != nil {
				log.Printf("Listen: error close conn %s", err.Error())
			}
		}()
		var incomeMsg model.InMessage
		err := websocket.JSON.Receive(wsc, &incomeMsg)
		if err != nil {
			fmt.Printf("failed receive json %s\n", err.Error())
			return
		}

		cl := NewClient(wsc, i, incomeMsg.UserID, incomeMsg.ChatID, incomeMsg.UserName)

		i.Add(cl)
		cl.Listen()
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		s := websocket.Server{
			Handler: onConnected,
		}
		s.ServeHTTP(w, r)
	})

	for {
		select {
		case c := <-i.addCh:
			i.ConnectChat(c)
			log.Println("Now", len(i.chats[c.ChatID].streams), "clients connected to", c.ChatID)
		case msg := <-i.getMsgFromClientCh:
			i.SendMessageToClient(msg)
		case c := <-i.delCh:
			fmt.Printf("delCh c: %v\n", c)
			i.deleteChat(c)
		case err := <-i.errCh:
			fmt.Printf("errCh err: %v\n", err)
		case <-i.doneCh:
			fmt.Println("doneCh done")
			return
		}
	}
}

func (i *Implementation) Add(c *Client) {
	i.addCh <- c
}
func (i *Implementation) Del(c *Client) {
	i.delCh <- c
}
func (i *Implementation) GetMessageFromClient(msg *model.InMessage) {
	i.getMsgFromClientCh <- msg
	for {
		i.mxChannel.Lock()
		_, ok := i.channels[msg.ChatID]
		i.mxChannel.Unlock()
		if !ok {
			continue
		} else {
			i.channels[msg.ChatID] <- msg
			return
		}
	}

}
func (i *Implementation) Done() {
	i.doneCh <- true
}
func (i *Implementation) Err(err error) {
	i.errCh <- err
}

func (i *Implementation) deleteChat(c *Client) {
	i.chats[string(c.ChatID)].m.Lock()
	delete(i.chats[string(c.ChatID)].streams, c.ID)
	i.chats[string(c.ChatID)].m.Unlock()
}
