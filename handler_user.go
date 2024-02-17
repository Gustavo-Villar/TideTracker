package main

import (
	"encoding/json" // Used for encoding and decoding JSON.
	"fmt"           // Used for formatting strings.
	"net/http"      // Provides HTTP client and server implementations.
	"time"          // Provides functionality for measuring and displaying time.

	"github.com/Gustavo-Villar/TideTracker/internal/database" // Internal package for database operations.
	"github.com/google/uuid"                                  // Used for generating UUIDs.
)

// handlerCreateUser is an HTTP handler for creating a new user in the database.
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// Define a struct to parse the request body.
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	paramns := parameters{}
	err := decoder.Decode(&paramns)
	if err != nil {
		// Respond with an error if the JSON decoding fails.
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Create a new user in the database with the provided name.
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(), // Generate a new UUID for the user.
		Name:      paramns.Name,
		CreatedAt: time.Now(), // Set the current time as the creation time.
		UpdatedAt: time.Now(), // Set the current time as the update time.
	})
	if err != nil {
		// Respond with an error if creating the user fails.
		respondWithError(w, 500, fmt.Sprintf("Error Creating User: %v", err))
		return
	}
	// Respond with the created user in JSON format.
	respondWithJson(w, 201, databaseUserToUser(user))
}

// handlerGetUserByAPIKey is an HTTP handler that responds with the user associated with the provided API key.
func (apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	// Respond with the user in JSON format.
	respondWithJson(w, 200, databaseUserToUser(user))
}

// handlerGetPostsForUser is an HTTP handler that retrieves and responds with posts for the authenticated user.
func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// Retrieve posts for the user with a limit of 10 posts.
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		// Respond with an error if retrieving posts fails.
		respondWithError(w, 500, fmt.Sprintf("couldn't retrieve posts for user: %v", err))
		return
	}

	// Respond with the posts in JSON format.
	respondWithJson(w, 200, databasePostsToPosts(posts))
}
