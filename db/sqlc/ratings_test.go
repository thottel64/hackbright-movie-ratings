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
	movier, err := createRandomMovie(t)
	if err != nil {
		t.Errorf("could not create movie: %v", err)
	}
	userr, err := createRandomUser(t)
	if err != nil {
		t.Errorf("could not create random user: %v", err)
	}
	arg := CreateRatingParams{
		Score: sql.NullInt32{
			1, true,
		},
		MovieID: sql.NullInt32{
			movier.ID, true,
		},
		UserID: sql.NullInt32{
			userr.ID, true,
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
func TestGetMovie(t *testing.T) {
	movie1, err := createRandomMovie(t)
	if err != nil {
		t.Errorf("error creating random movie: %v", err)
	}
	movie2, err := TestQueries.GetMovie(context.Background(), movie1.ID)
	if err != nil {
		t.Errorf("error creating movie: %v", err)
	}
	if movie1.Title != movie2.Title {
		t.Errorf("test movies are not the same. Wanted %v, got %v", movie1.Title, movie2.Title)
	}
}
func TestUpdateRating(t *testing.T) {
	rating1, err := createRandomRating(t)
	if err != nil {
		t.Errorf("error creating random movie: %v", err)
	}
	oldscore := rating1.Score.Int32
	err = TestQueries.UpdateRating(context.Background(), UpdateRatingParams{
		Score: sql.NullInt32{
			Int32: 8,
			Valid: true,
		},
		ID: 1,
	})

	if rating1.Score.Int32 != oldscore {
		t.Errorf("error, score was not updated. Got %v, wanted %v", oldscore, rating1.Score.Int32)
	}
}
func TestListUser(t *testing.T) {
	list1, err := TestQueries.ListUser(context.Background(), ListUserParams{
		Limit:  1,
		Offset: 1,
	})
	if err != nil {
		t.Errorf("error listing users")
	}
	for _, user := range list1 {
		if user.Username == "" {
			t.Errorf("list failed. user list is empty")
		}
	}
}
func TestListMovies(t *testing.T) {
	list1, err := TestQueries.ListMovies(context.Background(), 1)
	if err != nil {
		t.Errorf("unable to get list of movies from testqueries.ListMovies")
	}
	for _, values := range list1 {
		if values.ID < 1 {
			t.Errorf("list failed. Movie list is empty")
		}
		if values.ID == int32(len(list1)) {
			break
		}
	}
}
func TestListRatingsByUser(t *testing.T) {
	list1, err := TestQueries.ListRatingsByUser(context.Background(), sql.NullInt32{
		Int32: 2,
		Valid: true,
	})
	if err != nil {
		t.Errorf("error creating list: %v", err)
	}
	for _, rating := range list1 {
		if rating.MovieID.Int32 == 0 {
			t.Errorf("error generating list. Wanted %v, got %v", rating.MovieID.Int32, 0)
		}
	}
}
func TestListRatingsByMovie(t *testing.T) {
	list1, err := TestQueries.ListRatingsByMovie(context.Background(), sql.NullInt32{
		Int32: 1,
		Valid: true,
	})
	if err != nil {
		t.Errorf("error creating list: %v", err)
	}
	for _, rating := range list1 {
		if rating.ID == 0 {
			t.Errorf("empty list")
		}
	}
}
func TestDeleteMovie(t *testing.T) {
	err := TestQueries.DeleteMovie(context.Background(), 1)
	if err != nil {
		t.Errorf("error deleting movie: %v", err)
	}
}
func TestDeleteRating(t *testing.T) {
	err := TestQueries.DeleteRating(context.Background(), 1)
	if err != nil {
		t.Errorf("error deleting movie: %v", err)
	}
}
func TestDeleteUser(t *testing.T) {
	err := TestQueries.DeleteUser(context.Background(), 1)
	if err != nil {
		t.Errorf("error deleting user: %v", err)
	}
}
