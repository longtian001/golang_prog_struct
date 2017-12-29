package handlers

import (
	"encoding/json"
	"net/http"
)

// api server status
type serverStatusResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

// func: URL handlers
func ServerStatusHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(serverStatusResponse{Status: "OK", Code: 200})
}
