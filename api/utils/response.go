package utils

import (
	"encoding/json"
	"net/http"
	"strings"
)

func SendApiResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

func SendApiErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	payload := map[string]interface{}{
		"error": capitalizeFirstLetter(err.Error()),
	}
	SendApiResponse(w, statusCode, payload)
}

func capitalizeFirstLetter(sentence string) string {
	if len(sentence) == 0 {
		return sentence
	}
	// capitalize the first letter and concatenate it with the rest of the sentence
	return strings.ToUpper(string(sentence[0])) + sentence[1:]
}
