-- +migrate Up
CREATE TABLE tokens (
    id SERIAL PRIMARY KEY,
    token VARCHAR(256) NOT NULL,
    is_refresh_token BOOLEAN NOT NULL DEFAULT FALSE,
    user_id INTEGER NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +migrate Down
DROP TABLE tokens;