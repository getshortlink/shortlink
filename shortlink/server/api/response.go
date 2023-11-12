package api

import (
	"encoding/json"
	"net/http"
)

// Respond writes the data to the ResponseWriter as JSON.
func Respond(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}
