package server

import (
	"encoding/json"
	"sync"
	"tictactoe-server/game"
	"tictactoe-server/message"

	"github.com/gorilla/websocket"
)

var GameManager = game.NewManager()

type Client struct {
	conn      *websocket.Conn //connection
	PlayerId  string
	send      chan []byte // send msg
	RoomID    string      // clients room id
	Index     int         // 0 and 1 depends on which player the client is
	closeOnce sync.Once
}

func (c *Client) Send(msg message.Message) {
	data, err := json.Marshal(msg)
	if err != nil {
		println("error in sendding data", err)
		return
	}
	c.send <- data
}

func (c *Client) Close() {
	c.closeOnce.Do(func() {
		close(c.send)
		c.conn.Close()
	})
}

func (c *Client) SetRoomID(Id string) {
	c.RoomID = Id
}

func (c *Client) GetRoomID() string {
	return c.RoomID
}
func (c *Client) SetIndex(i int) {
	c.Index = i
}
func (c *Client) GetIndex() int {
	return c.Index
}
