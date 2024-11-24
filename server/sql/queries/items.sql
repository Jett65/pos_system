-- name: CreateItem :one
INSERT INTO items (id, name, price, number_of, description)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetItem :one
SELECT * FROM items WHERE id=$1;

-- name: GetItems :many
SELECT * FROM items;

-- name: GetItemByName :many
SELECT * from items WHERE name=$1;

-- name: UpdateItem :one
UPDATE items SET name=$2, price=$3, number_of=$4, description=$5
WHERE id=$1
RETURNING *;

-- name: UpdateNumberOf :one
UPDATE items SET number_of=$2
WHERE id=$1
RETURNING *;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id=$1;
