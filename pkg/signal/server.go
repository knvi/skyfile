package signal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

// we need a signaling server to manage the WebRTC connections
// the signaling server will keep track of the devices in each community
// and the connections between devices

// server is inspired by the signaling server in the Pion WebRTC examples

type SignalServer struct {
	mu          sync.Mutex
	communities map[string][]string // communities is a map of community names to a slice of device IDs
	macs        map[string]string   // macs is a map of device IDs to MAC addresses
	connections map[string]net.Conn
	candidates  []string
}

func NewSignalServer() *SignalServer {
	return &SignalServer{
		communities: make(map[string][]string),
		macs:        make(map[string]string),
		connections: make(map[string]net.Conn),
		candidates:  []string{},
	}
}

func (s *SignalServer) HandleConnection(conn net.Conn) {
	// handle the connection

	go func() {
		for {
			// read a message from connection
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				// handle error
				// TODO: later
			}

			fmt.Println("Message received: ", message) // TODO: switch to logging

			values := make(map[string]json.RawMessage)

			// parse msg
			err = json.Unmarshal([]byte(message), &values)
			if err != nil {
				// handle error
				// TODO: later
			}

		}
	}()
}
