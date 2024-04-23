-- +migrate Up
CREATE TABLE users_groups (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);
-- +migrate Down
DROP TABLE users_groups;