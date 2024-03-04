package services

import (
	"netflix/models"
)

type Actor struct {
	Name string
}

func ListUniqueActors(credits []models.Credit) []Actor {

	uniqueActors := make(map[string]bool)

	// Iterate over credits and add actors to the map
	for _, credit := range credits {
		if credit.Role == "ACTOR" {
			uniqueActors[credit.Name] = true
		}
	}

	var actorList []Actor
	for actor := range uniqueActors {
		actorList = append(actorList, Actor{Name: actor})
	}

	return actorList

}
