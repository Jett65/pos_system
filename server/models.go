package main

import (
	"pos_system/internal/database"

	"github.com/google/uuid"
)

type Aisle struct {
    ID uuid.UUID `json:"id"`
    Name string `json:"name"`
}

func databaseAisleToAisle(dbAisle database.Aisle) Aisle {
    return Aisle{
        ID: dbAisle.ID,
        Name: dbAisle.Name,
    }
}

func databaseAislesToAisles(dbAisles []database.Aisle) []Aisle {
    aisles := []Aisle{}
    for _, dbdbAisles := range dbAisles {
        aisles = append(aisles, databaseAisleToAisle(dbdbAisles))
    }

    return aisles
}
