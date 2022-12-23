package handlers

import (
	"encoding/json"
	"net/http"
)

// List is model
type List struct {
	json.RawMessage
}

func SendData(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
