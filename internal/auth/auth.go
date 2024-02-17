package auth

import (
	"errors"   // Package for creating error values.
	"net/http" // Provides HTTP client and server implementations.
	"strings"  // Provides simple functions to manipulate UTF-8 encoded strings.
)

// GetAPIKey extracts an API Key from the headers of an HTTP Request.
// It expects the API Key to be provided in the 'Authorization' header in the format:
// Authorization: ApiKey <keyValue>
func GetAPIKey(headers http.Header) (string, error) {
	// Retrieve the value of the 'Authorization' header.
	val := headers.Get("Authorization")
	// If the header is not provided, return an error indicating that no auth key was provided.
	if val == "" {
		return "", errors.New("no auth key provided")
	}

	// Split the header value by space. The expected format is "ApiKey <keyValue>".
	vals := strings.Split(val, " ")
	// If the header does not consist of two parts, return an error indicating that the auth header is malformed.
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	// If the first part of the header is not "ApiKey", return an error indicating that the header format is incorrect.
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header. Does it start with 'ApiKey'?")
	}
	// If the header is correctly formatted, return the second part (the actual API key value) and nil for the error.
	return vals[1], nil
}
