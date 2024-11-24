package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"pos_system/internal/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

    _"github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	// DataBase
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	// Router
	router := chi.NewRouter()

    // Change the AllowdOrigins to the client url
    // Check what headers should be allowed
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	// Health and Error
	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	// Routes
    v1Router.Post("/aisles", apiCfg.handlerCreateAisle)
    v1Router.Get("/aisles", apiCfg.handlerGetAisles)
    v1Router.Get("/aisle/{aisleID}", apiCfg.handlerGetAisle)    
    v1Router.Delete("/aisle/{aisleID}", apiCfg.handlerDeleteAisle)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", portString),
	}

	log.Printf("Server running on port %s\n", portString)
	log.Fatal(srv.ListenAndServe())
}
