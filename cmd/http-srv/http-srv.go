package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/perocha/dapr-pocs/internal/httphandlers"
	"github.com/perocha/dapr-pocs/internal/logger"
)

func main() {
	l := logger.New("debug")

	l.Debug("Starting server...")
	router := mux.NewRouter()

	router.HandleFunc("/", httphandlers.MainHandler)
	router.HandleFunc("/hello", httphandlers.HelloHandler)
	router.HandleFunc("/hello/{name}", httphandlers.HelloHandler)

	http.Handle("/", router)

	AppPort := ":8080"
	l.Debug("Listening on port", AppPort)
	l.Info("Listening on port", AppPort)
	err := http.ListenAndServe("localhost"+AppPort, router)
	if err != nil {
		l.Debug("Error starting server:", err)
	}
}
