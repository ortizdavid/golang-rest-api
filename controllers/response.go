package controllers

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, message string, code int, data interface{}, count interface{}) {
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
	response := map[string]interface{}{
		"message": message,
		"status": code,
		"data": data,
		"count": count,
	}
	json.NewEncoder(w).Encode(response)
}

func WriteError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
	errResponse := map[string]interface{}{
		"error": message,
		"status": code,
	}
	json.NewEncoder(w).Encode(errResponse)
}