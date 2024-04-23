-- +migrate Up
CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL
);
-- +migrate Down
DROP TABLE groups;