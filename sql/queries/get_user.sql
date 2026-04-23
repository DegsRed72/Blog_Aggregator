-- name: GetUser :one
SELECT id, created_at, updated_at, name from users 
WHERE name = $1;

-- name: GetUserID :one
SELECT * from users
WHERE ID = $1;