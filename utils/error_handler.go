package utils

import (
	"encoding/json"
	"net/http"
)

// WriteErrorResponse writes a JSON error response with a status code.
func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
