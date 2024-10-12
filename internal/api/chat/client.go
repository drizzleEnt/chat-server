package chat

import (
	"fmt"
	"io"

	"github.com/drizzleent/chat-server/internal/model"
	"golang.org/x/net/websocket"
)

type Client struct {
	ws  *websocket.Conn
	srv *Implementation

	ID     int
	ChatID string
	Name   string
	chMsg  chan *model.OutMessage
	doneCh chan bool
}

func NewClient(ws *websocket.Conn, srv *Implementation, id int, chatId string, name string) *Client {
	return &Client{
		ws:     ws,
		srv:    srv,
		ID:     id,
		ChatID: chatId,
		Name:   name,
		chMsg:  make(chan *model.OutMessage, 50),
		doneCh: make(chan bool),
	}
}

func (c *Client) Conn() *websocket.Conn {
	return c.ws
}

func (c *Client) Done() {
	c.doneCh <- true
}

func (c *Client) Write(msg *model.OutMessage) {
	select {
	case c.chMsg <- msg:
	default:
		err := fmt.Errorf("client %d is disconected", c.ID)
		fmt.Println(err)
	}
}

func (c *Client) Listen() {
	go c.listenWrite()
	c.listenRead()
}

func (c *Client) listenWrite() {
	for {
		select {
		case msg := <-c.chMsg:
			err := websocket.JSON.Send(c.ws, msg)
			if err != nil {
				fmt.Println("error sending msg", err)
			}
		case <-c.doneCh:
			c.doneCh <- true
			return
		}
	}
}

func (c *Client) listenRead() {
	for {
		select {
		case <-c.doneCh:
			c.doneCh <- true
		default:
			var msg model.InMessage
			err := websocket.JSON.Receive(c.ws, &msg)
			if err == io.EOF {
				c.doneCh <- true
			} else if err != nil {
				c.srv.Err(err)
			} else {
				c.srv.GetMessageFromClient(&msg)
			}

		}
	}
}
