package game

import (
	"log"
	"tictactoe-server/message"
)

type Room struct {
	ID             string //same as RoomId for cleint
	Size           int
	Board          []int     // inside board it can be 0 1(X) 2(O)
	Players        [2]Player //two players in one room
	Turn           int       // 0 or 1 based on which players turn is in the room
	Over           bool      // its game over or not
	RematchRequest [2]bool   //rematch required both yes
}

// create new room , turn start at player 0
func NewRoom(id string, size int) *Room {
	return &Room{
		ID:    id,
		Size:  size,
		Board: make([]int, size*size),
		Turn:  0, //at first by default 0's player has furst chance
		Over:  false,
	}
}

// adding player to room
func (r *Room) AddPlayer(p Player, index int) {
	log.Println("during addplayer")
	log.Println("newplayer", index)
	r.Players[index] = p
	p.SetRoomID(r.ID)
	p.SetIndex(index)
}

func (r *Room) StartGame() {
	r.Turn = 0
	r.Over = false

	log.Println("Starting game")
	log.Println("start the game, player 0:", r.Players[0].GetIndex(), "player 1:", r.Players[1].GetIndex())

	for i, p := range r.Players {
		if p == nil {
			log.Println("ERROR: player", i, "is nil")
			return
		}

		p.Send(message.Message{
			Type:   string(message.Start),
			Player: i,
			Size:   r.Size,
		})
	}

	r.Broadcast(message.Message{
		Type:  string(message.Move),
		Board: r.Board,
		Turn:  r.Turn,
	})

	log.Println("START + initial MOVE sent")
}

func (r *Room) MakeMove(p Player, cell int) {
	log.Println("SERVER: MakeMove called by player", p.GetIndex(), "cell", cell)
	//is game over
	//check room exist
	//check players room
	//check valid cell apply move
	// check win
	//check draw
	//broadcast the meessage

	if r.Over == true {
		return
	}
	if p.GetRoomID() != r.ID {
		return
	}
	if p.GetIndex() != r.Turn {
		log.Println("SERVER: Wrong turn")
		return
	}
	if cell < 0 || cell >= r.Size*r.Size {
		return
	}
	if r.Board[cell] != 0 {
		log.Println("SERVER: Cell not empty")
		return
	}

	playerValue := p.GetIndex() + 1
	r.Board[cell] = playerValue

	//  Win check
	if CheckWin(r.Board, r.Size, playerValue) {
		r.Over = true
		r.Broadcast(message.Message{
			Type:   string(message.GameOver),
			Winner: p.GetIndex(),
			Board:  r.Board,
		})
		return
	}
	//  Draw check
	if CheckDraw(r.Board) {
		r.Over = true
		r.Broadcast(message.Message{
			Type:  string(message.GameDraw),
			Board: r.Board,
		})
		return
	}
	r.Turn = 1 - r.Turn
	log.Println("SERVER: Broadcasting MOVE, turn:", r.Turn)
	r.Broadcast(message.Message{
		Type:  string(message.Move),
		Board: r.Board,
		Turn:  r.Turn,
	})
}

// broadcast helper
func (r *Room) Broadcast(msg message.Message) {
	for _, p := range r.Players {
		if p != nil {
			p.Send(msg)
		}
	}
}
