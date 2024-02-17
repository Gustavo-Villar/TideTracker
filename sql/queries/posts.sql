-- name: CreatePost :one
INSERT INTO posts (id, title, description, url, feed_id, published_at, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) -- $1=id, $2=title, $3=description, $4=url, $5=feed_id, $6=published_at, $7=created_at, $8=updated_at
RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.* FROM posts
JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1  -- $1=user_id
ORDER BY posts.published_at DESC
LIMIT $2; -- $2=limit