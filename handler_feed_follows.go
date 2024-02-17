package main

import (
	"encoding/json" // Used for encoding and decoding JSON.
	"fmt"           // Used for formatted I/O, with functions similar to C's printf and scanf.
	"net/http"      // Provides HTTP client and server implementations.
	"time"          // Provides functionality for measuring and displaying time.

	"github.com/Gustavo-Villar/TideTracker/internal/database" // Internal package for database operations.
	"github.com/go-chi/chi/v5"                                // Chi router for HTTP routing.
	"github.com/google/uuid"                                  // Provides support for UUIDs, unique identifiers used for identifying resources.
)

// handlerCreateFeedFollow is an HTTP handler that allows an authenticated user to follow a feed by its ID.
func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	// Define a struct to parse the request body.
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"` // The UUID of the feed to follow.
	}
	decoder := json.NewDecoder(r.Body) // Initialize a new JSON decoder.

	params := parameters{}
	err := decoder.Decode(&params) // Decode the request body into the params struct.
	if err != nil {
		// Respond with an error if JSON decoding fails.
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Create a new feed follow record in the database.
	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(), // Generate a new UUID for the feed follow record.
		FeedID:    params.FeedID,
		UserID:    user.ID,    // Associate the feed follow record with the authenticated user.
		CreatedAt: time.Now(), // Set the current time as the creation time.
		UpdatedAt: time.Now(), // Set the current time as the update time.
	})
	if err != nil {
		// Respond with an error if creating the feed follow record fails.
		respondWithError(w, 500, fmt.Sprintf("Error Creating Feed Follow: %v", err))
		return
	}
	// Respond with the created feed follow record in JSON format.
	respondWithJson(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

// handlerGetFeedFollowsByUser retrieves and responds with all feed follow records for the authenticated user.
func (apiCfg *apiConfig) handlerGetFeedFollowsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// Retrieve all feed follow records for the user from the database.
	feed_follows, err := apiCfg.DB.GetFeedFollowsByUserId(r.Context(), user.ID)
	if err != nil {
		// Respond with an error if retrieving the feed follows fails.
		respondWithError(w, 500, fmt.Sprintf("Error Getting Feed Follows: %v", err))
		return
	}
	// Respond with the retrieved feed follow records in JSON format.
	respondWithJson(w, 200, databaseFeedFollowsToFeedFollows(feed_follows))
}

// handlerDeleteFeedFollow allows an authenticated user to delete a feed follow record by its ID.
func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	// Extract the feed follow ID from the URL parameter.
	feedFollowIDStr := chi.URLParam(r, "feedFollowId")
	feedFollowID, err := uuid.Parse(feedFollowIDStr) // Convert the string to a UUID.
	if err != nil {
		// Respond with an error if the UUID conversion fails.
		respondWithError(w, 400, fmt.Sprintf("Error converting UUID: %v", err))
		return
	}

	// Delete the feed follow record from the database.
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID, // Ensure the feed follow record is associated with the authenticated user.
	})
	if err != nil {
		// Respond with an error if deleting the feed follow record fails.
		respondWithError(w, 400, fmt.Sprintf("Error Deleting Feed Follow: %v", err))
		return
	}

	// Respond with an empty JSON object to indicate successful deletion.
	respondWithJson(w, 200, struct{}{})
}
