package handler

import (
	"github.com/gorilla/mux"
)

// Router is a main router called from main.go
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingHandler).Methods("GET").Name("Ping")
	return r
}
