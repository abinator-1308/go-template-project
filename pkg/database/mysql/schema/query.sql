-- name: GetUser :one
SELECT *
FROM users
WHERE id = ? and name = ?;

-- name: PersistUser :execresult
INSERT INTO users (id, name) VALUES (?, ?);
