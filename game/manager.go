package game

import (
	"sync"
	"tictactoe-server/message"

	"github.com/google/uuid"
)

type Manager struct {
	//Waiting *Room
	WaitingQueue []Player
	Rooms        map[string]*Room
	mu           sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		Rooms: make(map[string]*Room),
	}
}

func (m *Manager) Join(p Player) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.WaitingQueue = append(m.WaitingQueue, p)
	if len(m.WaitingQueue) >= 2 {
		p1 := m.WaitingQueue[0]
		p2 := m.WaitingQueue[1]

		//removing from queue
		m.WaitingQueue = m.WaitingQueue[2:]
		roomID := uuid.NewString()
		room := NewRoom(roomID, 3)

		room.AddPlayer(p1, 0)
		room.AddPlayer(p2, 1)

		m.Rooms[roomID] = room
		room.StartGame()

	} else {
		p.Send(message.Message{
			Type: string(message.Waiting),
		})
	}
	/*
		if m.Waiting == nil {
			RoomId := uuid.NewString()
			room := NewRoom(RoomId, 3) // (roomid , size)

			room.AddPlayer(p, 0)
			log.Println("after Addplayer")
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
	*/
}
