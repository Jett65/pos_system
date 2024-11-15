-- name: CreateAisles :one
INSERT INTO aisles (id, name)
VALUES (
    encode(sha256(random()::text::bytea), hex), 
    $1
) 
RETURNING *;

-- name: GetAisles :many
SELECT * FROM aisles WHERE id=$1;

-- name: DeleteAisles :exec
DELETE FROM aisles
WHERE id=$1;
