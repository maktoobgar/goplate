-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

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

-- name: GetUserWithTokenId :one
SELECT u.* FROM users u JOIN tokens t ON u.id = t.user_id WHERE t.id = $1;

-- name: UpdateMe :one
UPDATE users SET first_name = $1, last_name = $2, display_name = $3, gender = $4 WHERE id = $5 RETURNING *;

-- name: UpdateAvatar :one
UPDATE users SET avatar = $1 WHERE id = $2 RETURNING *;
