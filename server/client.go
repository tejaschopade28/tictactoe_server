package server

import (
	"encoding/json"
	"sync"
	"tictactoe-server/game"

	"github.com/gorilla/websocket"
)

var GameManager = game.NewManager()

type Client struct {
	conn      *websocket.Conn //connection
	send      chan []byte     // send msg
	RoomID    string          // clients room id
	Index     int             // 0 and 1 depends on which player the client is
	closeOnce sync.Once
}

func (c *Client) Send(msg any) {
	data, err := json.Marshal(msg)
	if err != nil {
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
