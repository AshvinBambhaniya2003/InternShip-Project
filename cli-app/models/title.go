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
