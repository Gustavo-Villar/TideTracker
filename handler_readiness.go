package main

import (
	"net/http"
	"os"
	"time"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	versionString := os.Getenv("VERSION")

	responseStruct := struct {
		Status  string    `json:"status"`
		Version string    `json:"version"`
		Time    time.Time `json:"time"`
	}{
		Status:  "OK",
		Version: versionString,
		Time:    time.Now(),
	}
	respondWithJson(w, 200, responseStruct)
}
