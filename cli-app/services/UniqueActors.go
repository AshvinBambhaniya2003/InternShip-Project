package services

import (
	"netflix/models"
)

func ListUniqueActors(credits []models.Credit) []string {

	uniqueActors := make(map[string]bool)

	// Iterate over credits and add actors to the map
	for _, credit := range credits {
		if credit.Role == "ACTOR" {
			uniqueActors[credit.Name] = true
		}
	}

	// Extract unique actors from the map
	var actors []string
	for actor := range uniqueActors {
		actors = append(actors, actor)
	}

	return actors

}