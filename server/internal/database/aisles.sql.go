// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: aisles.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createAisle = `-- name: CreateAisle :one
INSERT INTO aisles (id, name)
VALUES ($1, $2) 
RETURNING id, name
`

type CreateAisleParams struct {
	ID   uuid.UUID
	Name string
}

func (q *Queries) CreateAisle(ctx context.Context, arg CreateAisleParams) (Aisle, error) {
	row := q.db.QueryRowContext(ctx, createAisle, arg.ID, arg.Name)
	var i Aisle
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteAisles = `-- name: DeleteAisles :exec
DELETE FROM aisles
WHERE id=$1
`

func (q *Queries) DeleteAisles(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAisles, id)
	return err
}

const getAisle = `-- name: GetAisle :one
SELECT id, name FROM aisles WHERE id=$1
`

func (q *Queries) GetAisle(ctx context.Context, id uuid.UUID) (Aisle, error) {
	row := q.db.QueryRowContext(ctx, getAisle, id)
	var i Aisle
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getAisles = `-- name: GetAisles :many
SELECT id, name FROM aisles
`

func (q *Queries) GetAisles(ctx context.Context) ([]Aisle, error) {
	rows, err := q.db.QueryContext(ctx, getAisles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Aisle
	for rows.Next() {
		var i Aisle
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
