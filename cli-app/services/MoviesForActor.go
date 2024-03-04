package services

import (
	"netflix/models"
	"strings"
)

type ActorMovies struct {
	Title         string
	CharacterName string
}

func ListMoviesForActor(titles []models.Title, credits []models.Credit, actorName string) []ActorMovies {
	var actorMovies []ActorMovies

	for _, credit := range credits {
		if strings.EqualFold(credit.Name, actorName) {
			title := models.FindMovie(titles, credit.TitleID)
			if title != nil {
				actorMovies = append(actorMovies, ActorMovies{Title: title.Title, CharacterName: credit.Character})
			}
		}
	}
	return actorMovies
}
