package services

import (
	"netflix/models"
)

type TitleWithCredit struct {
	models.Title
	Credits []models.Credit
}

func ListTitlesWithCredits(titles []models.Title, credits []models.Credit) []TitleWithCredit {

	// Create a map to store credits by title ID
	creditMap := make(map[string][]models.Credit)
	for _, credit := range credits {
		creditMap[credit.TitleID] = append(creditMap[credit.TitleID], credit)
	}

	// Create a slice to store detailed titles
	var titleWithCredits []TitleWithCredit

	// Iterate over titles and populate detailedTitles
	for _, title := range titles {
		titleWithCredit := TitleWithCredit{Title: title, Credits: creditMap[title.ID]}
		titleWithCredits = append(titleWithCredits, titleWithCredit)
	}

	return titleWithCredits
}