-- name: GetPosts :one
SELECT * FROM posts
WHERE id = ? LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY created_at;

-- name: CreatePosts :one
INSERT INTO posts (
  title, content, created_at
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdatePosts :one
UPDATE posts
set title = ?,
content = ?
WHERE id = ?
RETURNING *;

-- name: DeletePosts :one
DELETE FROM posts
WHERE id = ?
RETURNING *;