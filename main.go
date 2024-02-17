package main

import (
	"database/sql" // Standard interface for SQL database operations.
	"log"          // Logging package for logging messages and errors.
	"net/http"     // Package for HTTP server and client functionalities.
	"os"           // Package for interacting with operating system functionalities, such as environment variables.
	"time"         // Package for working with dates and times.

	"github.com/Gustavo-Villar/TideTracker/internal/database" // Internal package for database query handling.
	"github.com/go-chi/chi/v5"                                // Chi router, a lightweight, idiomatic HTTP router for building Go HTTP services.
	"github.com/go-chi/cors"                                  // CORS middleware for Chi router.
	"github.com/joho/godotenv"                                // Package for loading environment variables from a .env file.

	_ "github.com/lib/pq" // PostgreSQL driver, used implicitly by database/sql.
)

// apiConfig holds application-wide configurations, including the database query interface.
type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env") // Load environment variables from .env file.

	portString := os.Getenv("PORT") // Retrieve the port number from environment variables.
	if portString == "" {
		log.Fatal("PORT is not found in the environment") // Exit if PORT is not set.
	}

	dbURL := os.Getenv("DB_URL") // Retrieve the database URL from environment variables.
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment") // Exit if DB_URL is not set.
	}

	conn, err := sql.Open("postgres", dbURL) // Open a connection to the PostgreSQL database.
	if err != nil {
		log.Fatal("Can't connect to database", err) // Exit if the database connection fails.
	}

	db := database.New(conn) // Initialize the database query interface.
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute) // Start the RSS feed scraping process in a background goroutine.

	router := chi.NewRouter() // Initialize a new Chi router.

	router.Use(cors.Handler(cors.Options{ // Configure CORS options.
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()                // Create a sub-router for version 1 of the API.
	v1Router.Get("/healthz", handlerReadiness) // Route for health check.
	v1Router.Get("/err", handlerErr)           // Route for triggering an error (likely for testing).

	// User routes
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUserByAPIKey))

	// Feed routes
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	// Feed follow routes
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollowsByUser))
	v1Router.Delete("/feed_follows/{feedFollowId}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	// Posts route
	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))

	router.Mount("/v1", v1Router) // Mount the version 1 sub-router onto the main router.

	srv := &http.Server{ // Configure the HTTP server.
		Handler: router,
		Addr:    ":" + portString, // Listen on the configured port.
	}

	log.Printf("Server starting on port %v", portString) // Log the server start message.
	err = srv.ListenAndServe()                           // Start the server.
	if err != nil {
		log.Fatal(err) // Exit if the server fails to start.
	}
}
