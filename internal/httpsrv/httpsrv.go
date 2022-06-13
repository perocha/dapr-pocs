package httpsrv

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/perocha/dapr-pocs/internal/logger"
)

// Log is the logger for the server.
var Log = logger.New("debug")

// MainHandler is the main handler for the server.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	Log.Debug("Hello, world.")
}

// HelloHandler is a handler for the /hello endpoint.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	Log.Debug("HelloHandler", name)
}

// Run starts the server.
func Run() {
	Log.Debug("Starting server...")
	router := mux.NewRouter()

	// Add handlers here.
	router.HandleFunc("/", MainHandler)
	router.HandleFunc("/hello", HelloHandler)
	router.HandleFunc("/hello/{name}", HelloHandler)

	http.Handle("/", router)

	AppPort := ":8080"
	Log.Debug("Listening on port", AppPort)
	Log.Info("Listening on port", AppPort)
	err := http.ListenAndServe("localhost"+AppPort, router)
	if err != nil {
		Log.Fatal("Error starting server:", err)
	}
}
