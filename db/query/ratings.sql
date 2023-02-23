-- name: CreateRating :one
INSERT INTO
    ratings (score, movie_id, user_username)
VALUES
    ($1, $2, $3) RETURNING *;

-- name: GetRating :one
SELECT
    *
FROM
    ratings
WHERE
        id = $1
    LIMIT
  1;

-- name: ListRatingsByMovie :many
SELECT
    *
FROM
    ratings
WHERE
        movie_id = $1
ORDER BY
    created_at DESC;


-- name: DeleteRating :exec
DELETE FROM
    ratings
WHERE
        id = $1;

-- name: UpdateRating :exec
UPDATE
    ratings
SET
    score = $1
WHERE
        id = $2 RETURNING *;

-- name: ListRatingsByUser :many
SELECT
    *
FROM
    ratings
WHERE
        user_username = $1
ORDER BY
    created_at DESC;
