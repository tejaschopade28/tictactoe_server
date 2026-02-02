package game

import "tictactoe-server/message"

type Player interface {
	Send(msg message.Message)
	GetRoomID() string
	SetRoomID(id string)
	GetIndex() int
	SetIndex(int)
}
