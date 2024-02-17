-- name: CreateUser :one
INSERT INTO users(id, name, created_at, updated_at, api_key)
VALUES ($1, $2, $3, $4,
 encode(sha256(random()::text::bytea), 'hex')
)   -- $1=id, $2=name, $3=created_at, $4=updated_at
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1; -- $1=api_key
