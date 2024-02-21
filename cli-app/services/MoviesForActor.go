package services

import (
	"netflix/models"
	"strings"
)

func ListMoviesForActor(titles []models.Title, credits []models.Credit, actorName string) map[string]string {
	actorMap := make(map[string]string)
	for _, credit := range credits {
		if strings.EqualFold(credit.Name, actorName) {
			title := models.FindMovie(titles, credit.TitleID)
			if title != nil {
				actorMap[title.Title] = credit.Character
			}
		}
	}

	return actorMap
}
