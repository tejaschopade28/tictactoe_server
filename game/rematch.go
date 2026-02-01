package game

import (
	"log"
	"tictactoe-server/message"
)

func (r *Room) RematchingRequest(p Player) {
	if !r.Over {
		return
	}
	player_index := p.GetIndex()
	r.RematchRequest[player_index] = true
	log.Println("REMATCH_REQUEST from player", player_index)

	r.Broadcast(message.Message{
		Type:     string(message.RematchUpdate),
		Accepted: r.RematchRequest,
	})
	if r.RematchRequest[0] && r.RematchRequest[1] {
		r.resetMatch()
	}
}

func (r *Room) resetMatch() {
	r.Board = make([]int, r.Size*r.Size)
	r.Turn = 0
	r.Over = false
	r.RematchRequest = [2]bool{false, false}

	r.Broadcast(message.Message{
		Type:  string(message.RematchStart),
		Board: r.Board,
		Turn:  r.Turn,
	})
}
