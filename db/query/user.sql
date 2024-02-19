-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email,
    password_changed_at,
    created_at
) VALUES (
 $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetUser0 :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUser3 :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET hashed_password = $2,
full_name = $3,
email = $4,
password_changed_at = $5,
created_at = $6
WHERE username = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;
