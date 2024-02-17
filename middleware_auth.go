package main

import (
	"fmt"      // Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	"net/http" // Package http provides HTTP client and server implementations.

	"github.com/Gustavo-Villar/TideTracker/internal/auth"     // Internal package for authentication functions.
	"github.com/Gustavo-Villar/TideTracker/internal/database" // Internal package for database interaction, particularly user-related operations.
)

// authedHandler defines a custom handler type that includes an additional parameter for the authenticated user.
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// middlewareAuth wraps a handler function with authentication middleware, ensuring that only requests with valid API keys can access it.
func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the API key from the request header.
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			// If the API key is missing or invalid, respond with a 403 Forbidden status.
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		// Use the extracted API key to retrieve the corresponding user from the database.
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		// Before checking the error, try to convert the database user to the application-level user struct.
		if databaseUserToUser(user) == (User{}) {
			// If the conversion results in an empty user struct, it means the user was not found.
			respondWithError(w, 404, "User Not Found")
			return
		}
		if err != nil {
			// If there was an error retrieving the user (other than not found), respond with a 500 Internal Server Error.
			respondWithError(w, 500, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		// If authentication succeeds, call the original handler function with the authenticated user.
		handler(w, r, user)
	}
}
