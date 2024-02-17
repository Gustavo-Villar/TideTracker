package main

import (
	"encoding/json" // Package for encoding and decoding JSON data.
	"log"           // Package for logging messages and errors.
	"net/http"      // Package for building HTTP servers and clients.
)

// respondWithError sends an error message with the specified HTTP status code, formatted as a JSON response.
func respondWithError(w http.ResponseWriter, code int, msg string) {
	// Log server-side errors (5XX) to help with debugging and monitoring.
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}

	// Define a local struct to format the error message in JSON.
	type errResponse struct {
		Error string `json:"error"` // The error message to be sent to the client.
	}

	// Use respondWithJson to send the error message in JSON format.
	respondWithJson(w, code, errResponse{
		Error: msg,
	})
}

// respondWithJson encodes the given payload as JSON and sends it to the client with the specified HTTP status code.
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	// Marshal the payload into JSON. If marshaling fails, log the error and send a 500 Internal Server Error response.
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v\n", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate the response is in JSON format.
	w.Header().Add("Content-Type", "application/json")
	// Write the HTTP status code to the response header.
	w.WriteHeader(code)
	// Write the JSON-encoded data as the response body.
	w.Write(dat)
}
