package chat

import (
	"sync/atomic"
	"time"
)

type Message struct {
	Username string
	Message  string
	Date     string
}

type Hub struct {
	Presence map[string]string
	History  []Message
	Counter  atomic.Uint64
}

func NewHub() *Hub {
	counter := atomic.Uint64{}
	counter.Add(1000)

	return &Hub{
		Presence: make(map[string]string),
		History:  make([]Message, 0),
		Counter:  counter,
	}
}

func (h *Hub) Subscribe(username string) {
	h.Presence[username] = time.Now().Format(time.TimeOnly)
}

func (h *Hub) NewMessage(username string, message string) Message {
	msg := Message{
		Username: username,
		Message:  message,
		Date:     time.Now().Format(time.TimeOnly),
	}
	h.History = append(h.History, msg)
	h.Presence[username] = time.Now().Format(time.TimeOnly)
	return msg
}

func (h *Hub) NextID() uint64 {
	return h.Counter.Add(1)
}
