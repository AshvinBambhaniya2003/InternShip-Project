package services

import (
	"netflix/models"
	"reflect"
	"testing"
)

func TestListMoviesForActor(t *testing.T) {
	// Sample data
	titles := []models.Title{
		{ID: "1", Title: "Movie 1", Type: "MOVIE"},
		{ID: "2", Title: "Movie 2", Type: "MOVIE"},
		{ID: "3", Title: "Show 1", Type: "SHOW"},
	}
	credits := []models.Credit{
		{PersonID: 1, TitleID: "1", Name: "Actor 1", Character: "Character 1", Role: "ACTOR"},
		{PersonID: 2, TitleID: "2", Name: "Actor 2", Character: "Character 2", Role: "ACTOR"},
		{PersonID: 3, TitleID: "3", Name: "Actor 1", Character: "Character 3", Role: "ACTOR"},
	}

	actorName := "Actor 1"

	expected := map[string]string{
		"Movie 1": "Character 1",
	}

	result := ListMoviesForActor(titles, credits, actorName)

	// Check if the result matches the expectation
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ListMoviesForActor() returned unexpected result:\nExpected: %v\nActual: %v", expected, result)
	}
}
