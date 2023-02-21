-- name: CreateUser :one
INSERT INTO users(
    username,
    password,
    email)
    values(
"thottel", "pword", "taylor.hottel@shipt.com"
) RETURNING *;

-- name: SelectUser :one
SELECT *
from users
WHERE users.username = "thottel" LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: DeleteUser :one
DELETE FROM users WHERE users.username = "thottel"
RETURNING *;