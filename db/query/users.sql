-- name: CreateUser :one
INSERT INTO
    users (username, password, email)
VALUES
    ($1, $2, $3) RETURNING *;

-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
        username = $1
    LIMIT
  1;


-- name: ListUser :many
SELECT * FROM users
                  LIMIT $1 OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;