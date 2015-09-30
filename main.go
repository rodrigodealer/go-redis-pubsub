package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigodealer/queue-event-dispatcher/handlers"
)

func main() {
	r := mux.NewRouter()
	go handlers.Subscribe()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/publish", handlers.PublishHandler)
	http.ListenAndServe(":8000", r)
}
