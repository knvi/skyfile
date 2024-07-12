package main

import "net/http"

func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()
	// Routes
	mux.HandleFunc("/", app.wsEndpoint)
	mux.HandleFunc("/getDevices", app.GetDevicesHandler)
	return mux
}
