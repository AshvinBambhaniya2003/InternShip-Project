package services

import (
	"netflix/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListUniqueActors(t *testing.T) {
	// Mock credits data
	credits := []models.Credit{
		{PersonID: 1, TitleID: "1", Name: "John Doe", Character: "Character 1", Role: "ACTOR"},
		{PersonID: 1, TitleID: "2", Name: "John Doe", Character: "Character 2", Role: "ACTOR"},
		{PersonID: 3, TitleID: "3", Name: "Jane Smith", Character: "Character 3", Role: "ACTOR"},
		{PersonID: 4, TitleID: "4", Name: "Jane Smith", Character: "Character 4", Role: "DIRECTOR"},
	}

	// Test case 1: Actors exist in credits
	uniqueActors := ListUniqueActors(credits)
	assert.Equal(t, []Actor{
		{Name: "John Doe"},
		{Name: "Jane Smith"},
	}, uniqueActors)

	// Test case 2: No actors in credits
	emptyCredits := []models.Credit{}
	uniqueActors = ListUniqueActors(emptyCredits)
	assert.Empty(t, uniqueActors)

	// Test case 3: Actors with different roles
	creditsWithDifferentRoles := []models.Credit{
		{PersonID: 1, TitleID: "1", Name: "John Doe", Character: "Character 1", Role: "ACTOR"},
		{PersonID: 2, TitleID: "2", Name: "John Doe", Character: "Character 2", Role: "DIRECTOR"},
		{PersonID: 3, TitleID: "3", Name: "Jane Smith", Character: "Character 3", Role: "PRODUCER"},
	}
	uniqueActors = ListUniqueActors(creditsWithDifferentRoles)
	assert.Equal(t, []Actor{
		{Name: "John Doe"},
	}, uniqueActors)
}
