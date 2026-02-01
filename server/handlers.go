package server

import "tictactoe-server/message"

type HandlerFunc func(c *Client, msg message.Message)

var handlers = map[message.Type]HandlerFunc{
	message.Join:           handleJoin,
	message.Move:           handleMove,
	message.Leave:          handleLeave,
	message.RematchRequest: handleRematch,
}

// handle the correspoinding function

func handleJoin(c *Client, msg message.Message) {
	GameManager.Join(c)
}

func handleMove(c *Client, msg message.Message) {

	if msg.Cell < 0 {
		return
	}
	GameManager.HandleMove(c, msg.Cell)
}

func handleLeave(c *Client, msg message.Message) {
	GameManager.HandleLeave(c)
}
func handleRematch(c *Client, msg message.Message) {
	GameManager.HandleRematch(c)
}
