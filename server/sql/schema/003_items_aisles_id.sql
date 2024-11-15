-- +goose Up
ALTER TABLE items ADD COLUMN location UUID REFERENCES aisles(id);

-- +goose Down
ALTER TABLE items DROP COLUMN location;
