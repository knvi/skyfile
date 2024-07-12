package signal

import (
	"context"
	"encoding/json"
	"sync"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// we need a signaling server to manage the WebRTC connections
// the signaling server will keep track of the devices in each community
// and the connections between devices

// server is inspired by the signaling server in the Pion WebRTC examples

type SignalServer struct {
	mu          sync.Mutex
	communities map[string][]string
	macs        map[string]bool
	connections map[string]websocket.Conn
}

func NewSignalServer() *SignalServer {
	return &SignalServer{
		communities: map[string][]string{},
		macs:        map[string]bool{},
		connections: map[string]websocket.Conn{},
	}
}

func (s *SignalServer) HandleConnection(conn websocket.Conn) {
	// handle the connection

	go func() {
		for {
			_, data, err := conn.Read(context.Background())
			if err != nil {
				// TODO: handle
			}

			var v Message

			if err := json.Unmarshal(data, &v); err != nil {
				// TODO: handle
			}

			// handle messages
			switch v.Opcode {
			case "application":
				var app Application
				if err := json.Unmarshal(data, &app); err != nil {
					// TODO: handle
				}

				if _, ok := s.macs[app.mac]; ok {
					// reject cause the mac is already in use

					if err := wsjson.Write(context.Background(), &conn, &Rejection{Message: Message{"rejection"}}); err != nil {
						// TODO: handle
					}
					break
				}

				s.connections[app.mac] = conn

				// check if comm exists and < 2 devices

			}
		}
	}()
}
