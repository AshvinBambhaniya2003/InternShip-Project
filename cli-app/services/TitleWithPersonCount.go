package services

import (
	"netflix/models"
	"strings"
)

type TitleCount struct {
	Title string
	Count int
}

func ListTitleWithPersonCount(titles []models.Title, credits []models.Credit, searchQuery string) []TitleCount {

	var searchedTitles []models.Title
	for _, title := range titles {
		if searchQuery == "" || strings.Contains(strings.ToLower(title.Title), strings.ToLower(searchQuery)) {
			searchedTitles = append(searchedTitles, title)
		}
	}

	var titleCounts []TitleCount

	for _, title := range searchedTitles {
		count := countPersonsForTitle(title.ID, credits)
		title_count := TitleCount{
			Title: title.Title,
			Count: count,
		}
		titleCounts = append(titleCounts, title_count)
	}

	return titleCounts

}

func countPersonsForTitle(titleID string, credits []models.Credit) int {
	count := 0
	for _, credit := range credits {
		if credit.TitleID == titleID {
			count++
		}
	}
	return count
}