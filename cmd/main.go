package main

import (
	"log"
	"net/http"
	"os"
)

type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	app := &Application{
		errorLog: log.New(os.Stderr, "ERROR\t", log.LstdFlags),
		infoLog:  log.New(os.Stdout, "INFO\t", log.LstdFlags),
	}

	srv := &http.Server{
		Addr:     ":8080",
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	app.infoLog.Println("Starting server on :8080")
	err := srv.ListenAndServe()
	app.errorLog.Fatal(err)
}
