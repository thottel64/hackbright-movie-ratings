-- name: CreateRating :one
INSERT INTO ratings(
                    score,
                    user_username,
                    movie_id

) values (
             10,
          "thottel",
          1
         ) RETURNING *;
-- name: GetRating :one
SELECT * FROM ratings
where ratings.movie_id = 1 LIMIT 1;

-- name: ListRatingsByUser :many
SELECT * FROM ratings WHERE ratings.user_username = "thottel";

-- name: ListRatingByMovie :many
SELECT * FROM ratings
where ratings.movie_id = 1;

-- name: UpdateRating :one
UPDATE ratings
SET ratings.score = 9
WHERE ratings.movie_id = 1
RETURNING *;

-- name: DeleteRating :one
DELETE from ratings where ratings.id = 1 RETURNING *;