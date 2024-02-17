package main

import (
	"net/http" // Provides HTTP client and server implementations.
	"os"       // Provides a platform-independent interface to operating system functionality.
	"time"     // Provides functionality for measuring and displaying time.
)

// handlerReadiness checks the application's readiness to handle requests and responds with the application's status, version, and the current time.
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	// Retrieve the application version from environment variables.
	versionString := os.Getenv("VERSION")

	// Define a response structure to include the status, version, and current time.
	responseStruct := struct {
		Status  string    `json:"status"`  // The readiness status of the application.
		Version string    `json:"version"` // The current version of the application from the environment variable.
		Time    time.Time `json:"time"`    // The current server time.
	}{
		Status:  "OK",          // Indicate that the application is ready.
		Version: versionString, // Set the application version.
		Time:    time.Now(),    // Set the current time.
	}

	// Respond with the readiness information in JSON format.
	respondWithJson(w, 200, responseStruct)
}
