-- name: CreateMovie :one
INSERT INTO movies(
title,
overview,
release_date,
poster_url
) values (
          "Star Wars",
          "Star Wars is an American epic space opera multimedia franchise created by George Lucas, which began with the eponymous 1977 film and quickly became a worldwide pop culture phenomenon.",
          1978-07-21,
          "https://m.media-amazon.com/images/I/A1wnJQFI82L.jpg"
         ) RETURNING *;

-- name: GetMovieDetails :one
SELECT * FROM movies WHERE movies.title = "Star Wars";

-- name: ListMovies :many
SELECT title FROM movies;

-- name: DeleteMovie :one
DELETE FROM movies WHERE movies.title = "Star Wars"
    RETURNING *;