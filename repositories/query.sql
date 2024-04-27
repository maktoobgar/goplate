-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: RegisterUser :one
INSERT INTO users (
  phone_number, email, password
) VALUES (
  @phone_number, @email, @password
)
RETURNING *;
