-- name: CreateFeedFollow :one
INSERT INTO feed_follows(id, user_id, feed_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5) -- $1=id, $2=user_id, $3=feed_id, $4=created_at, $5=updated_at
RETURNING *;

-- name: GetFeedFollowsByUserId :many
SELECT * FROM feed_follows WHERE user_id = $1; -- $1=user_id

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE id = $1 AND user_id = $2; -- $1=id, %2=user_id

