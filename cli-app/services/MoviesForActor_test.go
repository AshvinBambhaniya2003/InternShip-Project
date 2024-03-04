package services

import (
	"netflix/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListMoviesForActor(t *testing.T) {
	// Mock titles and credits data
	titles := []models.Title{
		{ID: "1", Title: "Movie 1", Type: "MOVIE"},
		{ID: "2", Title: "Movie 2", Type: "MOVIE"},
		{ID: "3", Title: "TV Show 1", Type: "SHOW"},
	}

	credits := []models.Credit{
		{PersonID: 1, TitleID: "1", Name: "Actor 1", Character: "Character 1", Role: "ACTOR"},
		{PersonID: 2, TitleID: "2", Name: "Actor 2", Character: "Character 2", Role: "ACTOR"},
		{PersonID: 3, TitleID: "3", Name: "Actor 1", Character: "Character 3", Role: "ACTOR"},
	}

	// Test case 1: Actor has movies
	actorMovies := ListMoviesForActor(titles, credits, "Actor 1")
	assert.Equal(t, []ActorMovies{
		{Title: "Movie 1", CharacterName: "Character 1"},
	}, actorMovies)

	// Test case 2: Actor has no movies
	actorMovies = ListMoviesForActor(titles, credits, "Jane Smith")
	assert.Empty(t, actorMovies)

	// Test case 3: Actor does not exist
	actorMovies = ListMoviesForActor(titles, credits, "Unknown Actor")
	assert.Empty(t, actorMovies)
}
