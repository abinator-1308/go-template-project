-- name: GetUser :one
SELECT *
FROM users
WHERE id = ? and property = ?;

-- name: PersistUser :execresult
INSERT INTO users (id, name, property) VALUES (?, ?, ?);
