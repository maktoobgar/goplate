-- +migrate Up
CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    permission_id BIGINT NOT NULL,
    name VARCHAR(32) NOT NULL,
    user_id INTEGER,
    group_id INTEGER,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);
-- +migrate Down
DROP TABLE permissions;