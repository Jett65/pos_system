-- +goose Up

CREATE TABLE aisles (
    id UUID PRIMARY KEY, 
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE aisles
