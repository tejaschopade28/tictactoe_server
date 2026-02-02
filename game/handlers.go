package game

import (
	"log"
	"tictactoe-server/message"
)

// handlemove
func (m *Manager) HandleMove(p Player, cell int) {

	if p.GetRoomID() == "" {
		return
	}
	//find room
	room, ok := m.Rooms[p.GetRoomID()]
	if !ok {

		return
	}
	//pass move to game logic
	room.MakeMove(p, cell)
}

func (m *Manager) HandleLeave(p Player) {
	m.mu.Lock()
	defer m.mu.Unlock()

	//remove from queue if present
	for i, wp := range m.WaitingQueue {
		if wp == p {
			m.WaitingQueue = append(
				m.WaitingQueue[:i],
				m.WaitingQueue[i+1:]...,
			)
			return
		}
	}
	roomId := p.GetRoomID()
	if roomId == "" {
		return
	}
	room, ok := m.Rooms[roomId]

	if !ok {
		return
	}
	log.Println("Server leave the room ", roomId)
	room.Broadcast(message.Message{
		Type: string(message.OpponentLeft),
	})
	delete(m.Rooms, roomId)
}

func (m *Manager) HandleRematch(p Player) {
	roomId := p.GetRoomID()

	if roomId == "" {
		return
	}
	room, ok := m.Rooms[roomId]
	if !ok {
		return
	}
	room.RematchingRequest(p)

}
