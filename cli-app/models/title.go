package models

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

type Title struct {
	ID                  string
	Title               string
	Type                string
	Description         string
	ReleaseYear         int
	AgeCertification    string
	Runtime             int
	Genres              []string
	ProductionCountries []string
	Seasons             int
	IMDbID              string
	IMDbScore           float64
	IMDbVotes           int
	TmdbPopularity      float64
	TmdbScore           float64
}

func ReadTitles(filename string) ([]Title, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var titles []Title
	first := true

	for _, record := range records {
		// Skip the header row
		if first {
			first = false
			continue
		}

		title := Title{
			ID:                  record[0],
			Title:               record[1],
			Type:                record[2],
			Description:         record[3],
			ReleaseYear:         parseYear(record[4]),
			AgeCertification:    record[5],
			Runtime:             parseYear(record[6]),
			Genres:              []string{record[7]},
			ProductionCountries: []string{record[8]},
			Seasons:             parseYear(record[9]),
			IMDbID:              record[10],
			IMDbScore:           parseFloat(record[11]),
			IMDbVotes:           parseInt(record[12]),
			TmdbPopularity:      parseFloat(record[13]),
			TmdbScore:           parseFloat(record[14]),
		}

		titles = append(titles, title)
	}

	return titles, nil
}

func parseYear(yearStr string) int {
	return parseInt(yearStr)
}

func parseFloat(value string) float64 {
	if value == "" {
		return 0.00
	} else {
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal(err)
		}

		return floatValue
	}
}

func parseInt(value string) int {
	if value == "" {
		return 0
	} else {
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal(err)
		}

		return int(floatValue)
	}
}

func FindMovie(titles []Title, id string) *Title {
	for _, title := range titles {
		if title.ID == id && strings.EqualFold(title.Type, "MOVIE") {
			return &title
		}
	}
	return nil
}

func ListMoviesCountByReleaseYear(titles []Title) map[int]int {
	moviesCountByReleaseYear := make(map[int]int)
	for _, record := range titles {
		if record.Type == "MOVIE" {
			moviesCountByReleaseYear[record.ReleaseYear]++
		}
	}

	return moviesCountByReleaseYear
}

func ListMoviesCountByAgeCertificate(titles []Title) map[string]int {
	moviesCountByAgeCertificate := make(map[string]int)
	for _, record := range titles {
		if record.Type == "MOVIE" {
			moviesCountByAgeCertificate[record.AgeCertification]++
		}
	}

	return moviesCountByAgeCertificate
}

func ListMovieCountByRuntime(titles []Title) map[int]int {

	titlesCountByRuntimeMap := make(map[int]int)

	threshold := []int{30, 45, 60, 120, 360, 1000}

	for i := 0; i < len(threshold); i++ {
		titlesCountByRuntime := GetRuntimeWiseMovieCount(titles, threshold[i])
		titlesCountByRuntimeMap[threshold[i]] = titlesCountByRuntime
	}

	return titlesCountByRuntimeMap
}

func GetRuntimeWiseMovieCount(titles []Title, threshold int) int {
	count := 0
	for _, title := range titles {
		if title.Runtime < threshold && title.Type == "MOVIE" {
			count++
		}
	}
	return count
}

func ListTitlesCountPercentageByGenres(titles []Title) (map[string]int, int) {

	genreCount := make(map[string]int)

	// Count the occurrences of each genre
	totalTitles := 0
	for _, record := range titles {
		genres := strings.Split(record.Genres[0], ",")
		for _, genre := range genres {
			genre = strings.Trim(genre, "[]'\" ")
			genreCount[genre]++
			totalTitles++
		}
	}

	return genreCount, totalTitles
}

func ListTitlesCountPercentageByCountry(titles []Title) (map[string]int, int) {

	countryCount := make(map[string]int)

	// Count the occurrences of each genre
	totalTitles := 0
	for _, record := range titles {
		countries := strings.Split(record.ProductionCountries[0], ",")
		for _, country := range countries {
			country = strings.Trim(country, "[]'\" ")
			countryCount[country]++
			totalTitles++
		}
	}

	return countryCount, totalTitles
}

func ListTitleCountBySeasons(titles []Title) map[int]int {

	seasonsCounts := make(map[int]int)

	for _, record := range titles {
		if record.Seasons != 0 {
			seasonsCounts[record.Seasons]++
		}
	}

	return seasonsCounts
}

func ListTitlesCountByIMDbScore(titles []Title) map[[2]int]int {

	titlesCountByImdbMap := make(map[[2]int]int)

	threshold := [][2]int{{0, 4}, {4, 6}, {6, 8}, {8, 10}}

	for i := 0; i < len(threshold); i++ {
		titlesCountByImdb := GetImdbWiseTitleCount(titles, threshold[i][0], threshold[i][1])
		titlesCountByImdbMap[threshold[i]] = titlesCountByImdb
	}

	return titlesCountByImdbMap
}

func GetImdbWiseTitleCount(titles []Title, minThreshold int, maxThreshold int) int {
	count := 0
	for _, title := range titles {
		if title.IMDbScore >= float64(minThreshold) && title.IMDbScore < float64(maxThreshold) {
			count++
		}
	}
	return count
}

func ListTitlesCountByRuntime(titles []Title) map[int]int {

	titlesCountByRuntimeMap := make(map[int]int)

	threshold := []int{0, 30, 45, 60, 120, 360}

	for i := 0; i < len(threshold); i++ {
		titlesCountByRuntime := GetRuntimeWiseTitleCount(titles, threshold[i])
		titlesCountByRuntimeMap[threshold[i]] = titlesCountByRuntime
	}

	return titlesCountByRuntimeMap
}

func GetRuntimeWiseTitleCount(titles []Title, threshold int) int {
	count := 0
	for _, title := range titles {
		if title.Runtime > threshold {
			count++
		}
	}
	return count
}

func GetTitleTypeCountsAndPercentages(titles []Title) (int, float64, int, float64) {

	movieCount := 0
	showCount := 0

	for _, record := range titles {
		// Check the title type (movie or show)
		if record.Type == "MOVIE" {
			movieCount++
		} else if record.Type == "SHOW" {
			showCount++
		}
	}

	// Calculate total count and percentages
	totalCount := movieCount + showCount
	moviePercentage := float64(movieCount) / float64(totalCount) * 100
	showPercentage := float64(showCount) / float64(totalCount) * 100

	return movieCount, moviePercentage, showCount, showPercentage
}
