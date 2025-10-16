package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"unicode"
)

type response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimSpace(r.URL.Query().Get("name"))
	if name == "" {
		writeError(w, http.StatusBadRequest, "Invalid Input")
		return
	}

	first := unicode.ToUpper(rune(name[0]))
	if first >= 'A' && first <= 'M' {
		writeJSON(w, http.StatusOK, response{Message: "Hello " + name})
	} else if first >= 'N' && first <= 'Z' {
		writeError(w, http.StatusBadRequest, "Invalid Input")
	} else {
		writeError(w, http.StatusBadRequest, "Invalid Input")
	}
}

func writeJSON(w http.ResponseWriter, status int, data response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, response{Error: msg})
}
