-- +goose Up

CREATE TABLE items (
    id UUID PRIMARY KEY, 
    name TEXT NOT NULL, 
    price NUMERIC(10,2) NOT NULL,
    number_of INTEGER NOT NULL,
    description TEXT NOT NULL
);


-- location UUID REFERENCES aisles(id)

-- +goose Down
DROP TABLE items
