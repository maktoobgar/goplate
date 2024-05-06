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

-- name: UpdateUserParams :one
UPDATE users SET params = $1 WHERE id = $2 RETURNING *;

-- name: ConfirmPhoneNumber :one
UPDATE users SET phone_number_verified = TRUE WHERE id = $1 RETURNING *;

-- name: ConfirmEmail :one
UPDATE users SET email_verified = TRUE WHERE id = $1 RETURNING *;

-- name: ListUsers :many
WITH filtered_users AS (
    SELECT *
    FROM users
    WHERE 
        (phone_number LIKE '%' || @search || '%' OR email LIKE '%' || @search || '%' OR first_name LIKE '%' || @search || '%' OR last_name LIKE '%' || @search || '%' OR display_name LIKE '%' || @search || '%')
        AND (sqlc.narg('phone_number_verified')::boolean IS NULL OR phone_number_verified = @phone_number_verified)
        AND (sqlc.narg('email_verified')::boolean IS NULL OR email_verified = @email_verified)
        AND (sqlc.narg('gender')::int IS NULL OR gender = @gender)
        AND (sqlc.narg('is_active')::boolean IS NULL OR is_active = @is_active)
        AND (sqlc.narg('is_admin')::boolean IS NULL OR is_admin = @is_admin)
),
total_count AS (
    SELECT COUNT(*) AS total_count
    FROM filtered_users
)
SELECT *, (SELECT total_count FROM total_count) AS total_count
FROM filtered_users
ORDER BY 
    CASE WHEN sqlc.arg('desc')::boolean THEN
        CASE sqlc.arg('order_by')::text
            WHEN 'id' THEN id::text
            WHEN 'first_name' THEN first_name
            WHEN 'last_name' THEN last_name
            WHEN 'display_name' THEN display_name
            WHEN 'phone_number' THEN phone_number
            WHEN 'email' THEN email
            WHEN 'gender' THEN gender::text
            WHEN 'is_active' THEN is_active::text
            WHEN 'is_admin' THEN is_admin::text
            WHEN 'created_at' THEN created_at::text
            ELSE NULL
        END
    ELSE
        NULL
    END
    DESC,
    CASE WHEN NOT sqlc.arg('desc')::boolean THEN
        CASE sqlc.arg('order_by')::text
            WHEN 'id' THEN id::text
            WHEN 'first_name' THEN first_name
            WHEN 'last_name' THEN last_name
            WHEN 'display_name' THEN display_name
            WHEN 'phone_number' THEN phone_number
            WHEN 'email' THEN email
            WHEN 'gender' THEN gender::text
            WHEN 'is_active' THEN is_active::text
            WHEN 'is_admin' THEN is_admin::text
            WHEN 'created_at' THEN created_at::text
            ELSE NULL
        END
    ELSE
        NULL
    END
    ASC
LIMIT sqlc.arg('per_page')::int OFFSET ((sqlc.arg('page')::int - 1) * sqlc.arg('per_page')::int);
