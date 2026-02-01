package game

import (
	"log"
	"sync"
	"tictactoe-server/message"

	"github.com/google/uuid"
)

type Manager struct {
	Waiting *Room
	Rooms   map[string]*Room
	mu      sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		Rooms: make(map[string]*Room),
	}
}

func (m *Manager) Join(p Player) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.Waiting == nil {
		RoomId := uuid.NewString()
		room := NewRoom(RoomId, 3) // (roomid , size)
		room.AddPlayer(p, 0)
		m.Waiting = room
		m.Rooms[room.ID] = room
		p.Send(message.Message{
			Type: string(message.Waiting),
		})
		log.Println("room has one member")

	} else {
		log.Println("another member is joined in rooom")
		room := m.Waiting
		room.AddPlayer(p, 1)
		m.Waiting = nil
		room.StartGame()
		log.Println("another member is joined")
	}
}
