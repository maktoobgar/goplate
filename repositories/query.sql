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

-- name: CreateToken :one
INSERT INTO tokens (
  user_id, created_at
) VALUES (
  @user_id, @created_at
)
RETURNING *;

-- name: GetToken :one
SELECT * FROM tokens WHERE id = $1;

-- name: GetUserWithTokenId :one
SELECT u.* FROM users u JOIN tokens t ON u.id = t.user_id WHERE t.id = $1;
