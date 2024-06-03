package websocket

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		HandshakeTimeout: time.Second * 10,
	}

	sockets sync.Map

	joinChannel = make(chan *Socket)
	quitChannel = make(chan *Socket)
)

type Socket struct {
	sync.Mutex
	conn *websocket.Conn

	ID          string `json:"id"`
	BroadcastId string `json:"broadcastId"`
	Timestamp   int64  `json:"timestamp"`

	Tag map[string]interface{} `json:"tag"`
}

func NewSocket(conn *websocket.Conn, id, broadcastId string, tag map[string]interface{}) *Socket {
	return &Socket{
		conn:        conn,
		ID:          id,
		BroadcastId: broadcastId,
		Timestamp:   time.Now().Unix(),
		Tag:         tag,
	}
}

func (s *Socket) WriteMessage(message []byte) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.conn.WriteMessage(websocket.BinaryMessage, message)
}

func (s *Socket) BroadcastMessage(message []byte) error {
	if

	m := make(map[string]interface{})
	if err := json.Unmarshal(message, &m); err != nil {
		return err
	}
	m["broadcastId"] = s.BroadcastId
	message, _ = json.Marshal(&m)
	return s.WriteMessage(message)
}
