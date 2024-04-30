-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: RegisterUser :one
INSERT INTO users (
  phone_number, email, display_name, password, created_at
) VALUES (
  @phone_number, @email, @display_name, @password, @created_at
)
RETURNING *;
