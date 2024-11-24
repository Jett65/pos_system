package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pos_system/internal/database"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateAisle(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing JSON: %e", err))
        return
	}

	aisle, err := apiCfg.DB.CreateAisle(r.Context(), database.CreateAisleParams{
		ID:   uuid.New(),
		Name: params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create aisle: %e", err))
        return
	}

	respondWithJson(w, 201, databaseAisleToAisle(aisle))
}

func (apiCfg *apiConfig) handlerGetAisles(w http.ResponseWriter, r *http.Request) {
    aisle, err := apiCfg.DB.GetAisles(r.Context())
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Error Finding Aisles: %e", err))
        return
    }

    respondWithJson(w, 200, databaseAislesToAisles(aisle))    
}

func (apiCfg *apiConfig) handlerGetAisle(w http.ResponseWriter, r *http.Request) {
    aisleIDStr := chi.URLParam(r, "aisleID")

    aisleID, err := uuid.Parse(aisleIDStr)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Error Finding Aisle: %e", err))
        return
    }

    aisle, err := apiCfg.DB.GetAisle(r.Context(), aisleID)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Error Finding Aisle: %e", err))
        return
    }

    respondWithJson(w, 200, databaseAisleToAisle(aisle))
}

func (apiCfg *apiConfig) handlerDeleteAisle(w http.ResponseWriter, r *http.Request) { 
    aisleIDStr := chi.URLParam(r, "aisleID")

    aisleID, err := uuid.Parse(aisleIDStr)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Error Deleting Aisle: %e", err))
        return
    }

    err = apiCfg.DB.DeleteAisles(r.Context(), aisleID)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Error Deleting Aisle: %e", err))
        return
    }

    respondWithJson(w, 200, struct{}{})
}
