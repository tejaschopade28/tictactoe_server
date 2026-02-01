package message

type Type string

// client to server
const (
	Join           Type = "JOIN"
	Move           Type = "MOVE"
	Leave          Type = "LEAVE"
	RematchRequest Type = "REMATCH_REQUEST"
)

// server to client
const (
	Start         Type = "START"
	Waiting       Type = "WAITING"
	MoveUpdate    Type = "MOVE"
	GameOver      Type = "GAME_OVER"
	GameDraw      Type = "GAME_DRAW"
	RematchUpdate Type = "REMATCH_UPDATE"
	RematchStart  Type = "REMATCH_START"
	OpponentLeft  Type = "OPPONENT_LEFT"
)
