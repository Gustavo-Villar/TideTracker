package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Gustavo-Villar/TideTracker/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	paramns := parameters{}
	err := decoder.Decode(&paramns)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      paramns.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error Creating User: %v", err))
		return
	}
	respondWithJson(w, 200, databaseUserToUser(user))
}
