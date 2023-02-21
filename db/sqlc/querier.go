// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
)

type Querier interface {
	CreateMovie(ctx context.Context) (Movie, error)
	CreateRating(ctx context.Context) (Rating, error)
	CreateUser(ctx context.Context) (User, error)
	DeleteMovie(ctx context.Context) (Movie, error)
	DeleteRating(ctx context.Context) (Rating, error)
	DeleteUser(ctx context.Context) (User, error)
	GetMovieDetails(ctx context.Context) (Movie, error)
	GetRating(ctx context.Context) (Rating, error)
	ListMovies(ctx context.Context) ([]string, error)
	ListRatingByMovie(ctx context.Context) ([]Rating, error)
	ListRatingsByUser(ctx context.Context) ([]Rating, error)
	ListUsers(ctx context.Context) ([]User, error)
	SelectUser(ctx context.Context) (User, error)
	UpdateRating(ctx context.Context) (Rating, error)
}

var _ Querier = (*Queries)(nil)