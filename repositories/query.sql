-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: RegisterUser :one
INSERT INTO users (
  phone_number, email, display_name, password, created_at
) VALUES (
  @phone_number, @email, @display_name, @password, @created_at
)
RETURNING *;

-- name: LoginUserWithPhoneNumber :one
SELECT * FROM users WHERE phone_number = $1;

-- name: LoginUserWithEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateAccessToken :one
INSERT INTO tokens (
  token, is_refresh_token, user_id, expires_at, created_at
) VALUES (
  @token, FALSE, @user_id, @expires_at, @created_at
)
RETURNING *;

-- name: CreateRefreshToken :one
INSERT INTO tokens (
  token, is_refresh_token, user_id, expires_at, created_at
) VALUES (
  @token, TRUE, @user_id, @expires_at, @created_at
)
RETURNING *;
