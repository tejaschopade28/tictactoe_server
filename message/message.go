package message

type Message struct {
	Type     string  `json:"type"`
	Board    []int   `json:"board,omitempty"`
	Turn     int     `json:"turn,omitempty"`
	Cell     int     `json:"cell,omitempty"`
	Winner   int     `json:"winner,omitempty"`
	Player   int     `json:"playerIndex,omitempty"`
	Size     int     `json:"size,omitempty"`
	Accepted [2]bool `json:"accepted,omitempty"`
}
