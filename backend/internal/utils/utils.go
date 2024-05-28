package utils

import (
    "encoding/json"
    "net/http"
)

// RespondWithError sends an error response with the given message and status code
func RespondWithError(w http.ResponseWriter, message string, statusCode int) {
    response := map[string]string{"error": message}
    jsonResponse(w, response, statusCode)
}

// RespondWithJSON sends a JSON response with the given data and status code
func RespondWithJSON(w http.ResponseWriter, data interface{}, statusCode int) {
    jsonResponse(w, data, statusCode)
}

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
    response, err := json.Marshal(data)
    if err != nil {
        // If unable to marshal JSON, respond with internal server error
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    w.Write(response)
}
