package util

import (
	"encoding/json"
	"net/http"
)

// WriteJSON handles marshaling the payload to JSON and writing the response back out to the client
func WriteJSON(w http.ResponseWriter, r *http.Request, status int, payload interface{}) {
	// All responses are JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Open up CORS to all hosts
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", r.Method)

	out, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(status)
	w.Write(out)
	return
}
