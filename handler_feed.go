package main

import (
	"encoding/json" // Used for encoding and decoding JSON.
	"fmt"           // Used for formatted I/O, with functions similar to C's printf and scanf.
	"net/http"      // Provides HTTP client and server implementations.
	"time"          // Provides functionality for measuring and displaying time.

	"github.com/Gustavo-Villar/TideTracker/internal/database" // Internal package for database operations.
	"github.com/google/uuid"                                  // Provides support for UUIDs, unique identifiers used for feed and user identification.
)

// handlerCreateFeed is an HTTP handler for creating a new feed associated with the authenticated user.
func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	// Define a struct to parse the request body.
	type parameters struct {
		Name string `json:"name"` // Name of the feed.
		Url  string `json:"url"`  // URL of the feed.
	}
	decoder := json.NewDecoder(r.Body) // Initialize a new JSON decoder.

	params := parameters{}
	err := decoder.Decode(&params) // Decode the request body into the params struct.
	if err != nil {
		// Respond with an error if JSON decoding fails.
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Create a new feed in the database with the provided name, URL, and associated user ID.
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(), // Generate a new UUID for the feed.
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,    // Associate the feed with the authenticated user.
		CreatedAt: time.Now(), // Set the current time as the creation time.
		UpdatedAt: time.Now(), // Set the current time as the update time.
	})
	if err != nil {
		// Respond with an error if creating the feed fails.
		respondWithError(w, 500, fmt.Sprintf("Error Creating Feed: %v", err))
		return
	}
	// Respond with the created feed in JSON format.
	respondWithJson(w, 201, databaseFeedToFeed(feed))
}

// handlerGetFeeds is an HTTP handler that retrieves and responds with all feeds from the database.
func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	// Retrieve all feeds from the database.
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		// Respond with an error if retrieving feeds fails.
		respondWithError(w, 500, fmt.Sprintf("Error Getting Feeds: %v", err))
		return
	}
	// Respond with the retrieved feeds in JSON format.
	respondWithJson(w, 200, databaseFeedsToFeeds(feeds))
}
