package services

import (
	"netflix/models"
	"reflect"
	"testing"
)

func TestListTitlesWithCredits(t *testing.T) {
	// Sample titles
	titles := []models.Title{
		{ID: "1", Title: "Title 1"},
		{ID: "2", Title: "Title 2"},
	}

	// Sample credits
	credits := []models.Credit{
		{PersonID: 1, TitleID: "1", Name: "Actor 1"},
		{PersonID: 2, TitleID: "1", Name: "Actor 2"},
		{PersonID: 3, TitleID: "2", Name: "Actor 3"},
	}

	// Expected output
	expected := []TitleWithCredit{
		{
			Title:   titles[0],
			Credits: []models.Credit{{PersonID: 1, TitleID: "1", Name: "Actor 1"}, {PersonID: 2, TitleID: "1", Name: "Actor 2"}},
		},
		{
			Title:   titles[1],
			Credits: []models.Credit{{PersonID: 3, TitleID: "2", Name: "Actor 3"}},
		},
	}

	result := ListTitlesWithCredits(titles, credits)

	// Check if the result matches the expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ListTitlesWithCredits() returned unexpected result:\nExpected: %v\nActual: %v", expected, result)
	}
}
