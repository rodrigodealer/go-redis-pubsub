package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	go Subscribe()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/publish", PublishHandler)
	http.ListenAndServe(":8000", r)
}
