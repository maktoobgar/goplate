-- +migrate Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(16) NOT NULL UNIQUE,
    email VARCHAR(64),
    password VARCHAR(256) NOT NULL,
    profile VARCHAR(256),
    first_name VARCHAR(128),
    last_name VARCHAR(128),
    display_name VARCHAR(128) NOT NULL,
    gender INTEGER NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT FALSE,
    registered BOOLEAN NOT NULL DEFAULT FALSE,
    deactivation_reason VARCHAR(256),
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
    otp_remaining_attempts INTEGER NOT NULL DEFAULT 0,
    otp_code INTEGER,
    otp_due_date TIMESTAMP,
    is_superuser BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL
);
-- +migrate Down
DROP TABLE users;