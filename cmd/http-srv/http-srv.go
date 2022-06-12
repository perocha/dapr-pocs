package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/perocha/dapr-pocs/internal/httphandlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", httphandlers.MainHandler)
	router.HandleFunc("/hello", httphandlers.HelloHandler)
	router.HandleFunc("/hello/{name}", httphandlers.HelloHandler)

	http.Handle("/", router)

	AppPort := ":8080"
	log.Printf("Listening on port %s", AppPort)
	log.Fatal(http.ListenAndServe("localhost"+AppPort, router))
}
