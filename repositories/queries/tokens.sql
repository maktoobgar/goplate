-- name: CreateToken :one
INSERT INTO tokens (
  user_id, created_at
) VALUES (
  @user_id, @created_at
)
RETURNING *;

-- name: GetToken :one
SELECT * FROM tokens WHERE id = $1;
