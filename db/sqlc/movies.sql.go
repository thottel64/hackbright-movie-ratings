// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: movies.sql

package db

import (
	"context"
)

const createMovie = `-- name: CreateMovie :one
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
         ) RETURNING id, title, overview, release_date, poster_url
`

func (q *Queries) CreateMovie(ctx context.Context) (Movie, error) {
	row := q.db.QueryRowContext(ctx, createMovie)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Overview,
		&i.ReleaseDate,
		&i.PosterUrl,
	)
	return i, err
}

const deleteMovie = `-- name: DeleteMovie :one
DELETE FROM movies WHERE movies.title = "Star Wars"
    RETURNING id, title, overview, release_date, poster_url
`

func (q *Queries) DeleteMovie(ctx context.Context) (Movie, error) {
	row := q.db.QueryRowContext(ctx, deleteMovie)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Overview,
		&i.ReleaseDate,
		&i.PosterUrl,
	)
	return i, err
}

const getMovieDetails = `-- name: GetMovieDetails :one
SELECT id, title, overview, release_date, poster_url FROM movies WHERE movies.title = "Star Wars"
`

func (q *Queries) GetMovieDetails(ctx context.Context) (Movie, error) {
	row := q.db.QueryRowContext(ctx, getMovieDetails)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Overview,
		&i.ReleaseDate,
		&i.PosterUrl,
	)
	return i, err
}

const listMovies = `-- name: ListMovies :many
SELECT title FROM movies
`

func (q *Queries) ListMovies(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			return nil, err
		}
		items = append(items, title)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}