package main

import (
	"fmt"
	"net/http"

	"github.com/Gustavo-Villar/TideTracker/internal/auth"
	"github.com/Gustavo-Villar/TideTracker/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if databaseUserToUser(user) == (User{}) {
			respondWithError(w, 404, "User Not Found")
			return

		}
		if err != nil {
			respondWithError(w, 500, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
