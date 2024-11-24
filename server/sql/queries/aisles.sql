-- name: CreateAisle :one
INSERT INTO aisles (id, name)
VALUES ($1, $2) 
RETURNING *;

-- name: GetAisle :one
SELECT * FROM aisles WHERE id=$1;

-- name: GetAisles :many 
SELECT * FROM aisles;

-- name: DeleteAisles :exec
DELETE FROM aisles
WHERE id=$1;
