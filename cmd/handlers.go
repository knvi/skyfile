package main

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	mu       sync.Mutex
	devices  = make(map[*websocket.Conn]bool)
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func (app *Application) wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// handle error
	}

	mu.Lock()
	devices[ws] = true
	mu.Unlock()

	go echo(ws)
}

func echo(ws *websocket.Conn) {
	defer func() {
		mu.Lock()
		delete(devices, ws)
		mu.Unlock()
		ws.Close()
	}()

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}

func (app *Application) GetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Here you can return the list of devices to the client
	// For example, you can write them to the response writer
	for device := range devices {
		_, _ = w.Write([]byte(device.RemoteAddr().String() + "\n"))
	}
}
