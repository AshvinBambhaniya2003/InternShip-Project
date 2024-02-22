package services

import (
	"netflix/models"
	"reflect"
	"testing"
)

func TestListUniqueActors(t *testing.T) {
	// Sample credits
	credits := []models.Credit{
		{PersonID: 1, TitleID: "1", Name: "Actor 1", Role: "ACTOR"},
		{PersonID: 2, TitleID: "1", Name: "Actor 2", Role: "ACTOR"},
		{PersonID: 3, TitleID: "1", Name: "Actor 3", Role: "DIRECTOR"},
		{PersonID: 4, TitleID: "1", Name: "Actor 1", Role: "ACTOR"},
		{PersonID: 5, TitleID: "1", Name: "Actor 4", Role: "ACTOR"},
	}

	// Expected unique actors
	expected := []string{"Actor 1", "Actor 2", "Actor 4"}

	// Call ListUniqueActors function
	result := ListUniqueActors(credits)

	// Check if the result matches the expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ListUniqueActors() returned unexpected result:\nExpected: %v\nActual: %v", expected, result)
	}
}
