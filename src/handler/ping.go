package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PingRes is to return { "ping": "pong" }
type PingRes struct {
	Ping string `json:"ping"`
}

// PingHandler is a request handler called from router.go
func PingHandler(w http.ResponseWriter, req *http.Request) {
	resJSON, err := json.Marshal(PingRes{Ping: "pong"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(resJSON))
}
