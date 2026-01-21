// Package httputil provides HTTP response helper functions.
package httputil

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// JSON writes a JSON response.
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	}
}

// Text writes a plain text response.
func Text(w http.ResponseWriter, statusCode int, text string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(statusCode)
	if _, err := w.Write([]byte(text)); err != nil {
		slog.Error("failed to write response", "error", err)
	}
}
