package main

import (
	"fmt"
	"log"
	"os"
    "net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	// DataBase

    router := chi.NewRouter()

    // TODO: Set up cors

    v1Router := chi.NewRouter()

    // Health and Error
    v1Router.Get("/health", handlerReadiness)
    v1Router.Get("/err", handlerErr)

    // Routes

    router.Mount("/v1", v1Router)

    srv := &http.Server{
        Handler: router,
        Addr: fmt.Sprintf(":%s", portString),
    }

    log.Printf("Server running on port %s\n", portString)
    log.Fatal(srv.ListenAndServe())
}
