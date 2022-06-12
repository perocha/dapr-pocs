package httphandlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world.")
	log.Printf("Hello, world.")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Hello, %s.", name)
	log.Printf("Hello, %s.", name)
}
