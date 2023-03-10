-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM users
WHERE username = ? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO users (username, password, role)
VALUES (?, ?, ?);