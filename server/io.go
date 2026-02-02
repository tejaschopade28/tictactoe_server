package server

import (
	"encoding/json"
	"log"
	"tictactoe-server/message"

	"github.com/gorilla/websocket"
)

func (c *Client) reading() {

	defer func() {
		//c.hub.unregister <- c
		// can be recover from paniic reading
		if r := recover(); r != nil {
			log.Println("reading panic recoverd,", r)
		}
		c.Close()
		//close(c.send)
		log.Println("client disconnected")
	}()
	for {
		_, payload, err := c.conn.ReadMessage()

		if err != nil {
			log.Println("Read error!")
			break
		}

		var msg message.Message
		log.Printf("RAW MSG: %s\n", string(payload))
		if err := json.Unmarshal(payload, &msg); err != nil {
			log.Println("JSON unmarshal Error:", err)
			continue
		}

		log.Printf("Parsed Msg: %+v\n", msg)

		handler, ok := handlers[message.Type(msg.Type)]
		if !ok {
			log.Println("UNKNOWN TYPE:", msg.Type, " ", message.Type(msg.Type))
			continue
		}

		handler(c, msg)
	}
}
func (c *Client) writing() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("writing panic recoverd", r)
		}
		c.Close()
	}()
	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Error in wriiting", err)
			break
		}
	}
}
