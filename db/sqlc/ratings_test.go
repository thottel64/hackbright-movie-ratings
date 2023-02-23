package db

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"
	"time"
)

func createRandomRating(t *testing.T) (Rating, error) {
	rand.Seed(time.Now().Unix())
	arg := CreateRatingParams{
		Score: sql.NullInt32{
			Int32: int32(rand.Int()),
			Valid: true,
		},
	}

	rating, err := TestQueries.CreateRating(context.Background(), arg)
	if err != nil {
		t.Errorf("error creating rating: %v", err)
		return Rating{}, err
	}

	if rating.Score.Int32 == 0 {
		t.Errorf("got %v, wanted %v", rating.Score, arg.Score)
		return Rating{}, err
	}
	return rating, nil
}

func createRandomMovie(t *testing.T) (Movie, error) {
	rand.Seed(time.Now().Unix())

	arg := CreateMovieParams{
		Title:       "test-film",
		Overview:    "test-description",
		ReleaseDate: time.Now(),
		PosterUrl:   "https://media.istockphoto.com/id/1300341704/vector/tv-screen-test-television-test-pattern-stripes-retro-style-screensaver-vector-illustration.jpg?s=612x612&w=0&k=20&c=24Yn9zpQZjMM2O0CcG0x0yyNV2DcAnJvP5kN4syM2J4=",
	}
	movie, err := TestQueries.CreateMovie(context.Background(), arg)
	if movie.Title != arg.Title {
		t.Errorf("got %v, wanted %v", movie.Title, arg.Title)
		return Movie{}, err
	}
	if err != nil {
		t.Errorf("error creating movie:\n %v", err)
		return Movie{}, err
	}
	return movie, nil
}
func createRandomUser(t *testing.T) (User, error) {
	rand.Seed(time.Now().Unix())
	arg := CreateUserParams{
		Username: "test-user",
		Password: "test-pword",
		Email:    "test-email@something.com",
	}
	user, err := TestQueries.CreateUser(context.Background(), arg)
	if user.Username != arg.Username {
		t.Errorf("got %v, wanted %v", user.Username, arg.Username)
		return User{}, err
	}
	if err != nil {
		t.Errorf("error creating user: \n %v", err)
		return User{}, err
	}
	return user, nil
}
func TestCreateMovie(t *testing.T) {
	createRandomMovie(t)
}
func TestCreateRating(t *testing.T) {
	createRandomRating(t)
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
func TestGetRating(t *testing.T) {
	rating1, err := createRandomRating(t)
	if err != nil {
		t.Errorf("error creating random rating: %v", err)
	}
	rating2, err := TestQueries.GetRating(context.Background(), rating1.ID)
	if err != nil {
		t.Errorf("error creating rating: %v", err)
	}
	if rating1.Score != rating2.Score {
		t.Errorf("test ratings are not the same, %v", err)
	}
}
func TestGetUser(t *testing.T) {
	user1, err := createRandomUser(t)
	if err != nil {
		t.Errorf("error creating random user: %v", err)
	}
	user2, err := TestQueries.GetUser(context.Background(), user1.Username)
	if err != nil {
		t.Errorf("error creating user: %v", err)
	}
	if user2.Username != user1.Username {
		t.Errorf("test users are not the same. Wanted %v, got %v", user1.Username, user2.Username)
	}
}
