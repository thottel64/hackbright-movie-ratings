// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateMovie(ctx context.Context, arg CreateMovieParams) (Movie, error)
	CreateRating(ctx context.Context, arg CreateRatingParams) (Rating, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteMovie(ctx context.Context, id int32) error
	DeleteRating(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int32) error
	GetMovie(ctx context.Context, id int32) (Movie, error)
	GetRating(ctx context.Context, id int32) (Rating, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListMovies(ctx context.Context, id int32) ([]Movie, error)
	ListRatingsByMovie(ctx context.Context, movieID sql.NullInt32) ([]Rating, error)
	ListRatingsByUser(ctx context.Context, userID sql.NullInt32) ([]Rating, error)
	ListUser(ctx context.Context, arg ListUserParams) ([]User, error)
	UpdateRating(ctx context.Context, arg UpdateRatingParams) error
}

var _ Querier = (*Queries)(nil)
