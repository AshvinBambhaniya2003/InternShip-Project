package models

import (
	"math"
	"os"
	"reflect"
	"testing"
)

func TestReadTitles(t *testing.T) {
	// Create a temporary CSV file
	file, err := os.CreateTemp("", "test_titles.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Write test data to the CSV file
	data := `ID,Title,Type,Description,ReleaseYear,AgeCertification,Runtime,Genres,ProductionCountries,Seasons,IMDbID,IMDbScore,IMDbVotes,TmdbPopularity,TmdbScore
1,Title 1,Movie,Description 1,2020,PG-13,120,Action|Adventure|Drama,USA,1,tt1234567,7.5,100,65.4,7.3
2,Title 2,TV Show,Description 2,2019,TV-MA,45,Drama,UK,3,tt7654321,8.2,200,78.2,8.0
`
	_, err = file.WriteString(data)
	if err != nil {
		t.Fatal(err)
	}
	file.Close()

	// Test reading titles from the created CSV file
	titles, err := ReadTitles(file.Name())
	if err != nil {
		t.Errorf("Error reading titles: %v", err)
	}

	// Expected titles
	expected := []Title{
		{
			ID:                  "1",
			Title:               "Title 1",
			Type:                "Movie",
			Description:         "Description 1",
			ReleaseYear:         2020,
			AgeCertification:    "PG-13",
			Runtime:             120,
			Genres:              []string{"Action|Adventure|Drama"},
			ProductionCountries: []string{"USA"},
			Seasons:             1,
			IMDbID:              "tt1234567",
			IMDbScore:           7.5,
			IMDbVotes:           100,
			TmdbPopularity:      65.4,
			TmdbScore:           7.3,
		},
		{
			ID:                  "2",
			Title:               "Title 2",
			Type:                "TV Show",
			Description:         "Description 2",
			ReleaseYear:         2019,
			AgeCertification:    "TV-MA",
			Runtime:             45,
			Genres:              []string{"Drama"},
			ProductionCountries: []string{"UK"},
			Seasons:             3,
			IMDbID:              "tt7654321",
			IMDbScore:           8.2,
			IMDbVotes:           200,
			TmdbPopularity:      78.2,
			TmdbScore:           8.0,
		},
	}

	// Compare actual and expected titles
	if !reflect.DeepEqual(titles, expected) {
		t.Errorf("ReadTitles() returned unexpected titles:\nExpected: %v\nActual: %v", expected, titles)
	}
}

func TestListMoviesCountByReleaseYear(t *testing.T) {
	titles := []Title{
		{ID: "1", Title: "Movie 1", Type: "MOVIE", ReleaseYear: 2000},
		{ID: "2", Title: "Movie 2", Type: "MOVIE", ReleaseYear: 2000},
		{ID: "3", Title: "Show 1", Type: "SHOW", ReleaseYear: 2001},
		{ID: "4", Title: "Movie 3", Type: "MOVIE", ReleaseYear: 2002},
		{ID: "5", Title: "Movie 4", Type: "MOVIE", ReleaseYear: 2000},
		{ID: "6", Title: "Show 2", Type: "SHOW", ReleaseYear: 2001},
	}

	expected := map[int]int{
		2000: 3,
		2002: 1,
	}

	result := ListMoviesCountByReleaseYear(titles)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ListMoviesCountByReleaseYear() returned unexpected result:\nExpected: %v\nActual: %v", expected, result)
	}
}

func TestListMoviesCountByAgeCertificate(t *testing.T) {
	titles := []Title{
		{ID: "1", Title: "Movie 1", Type: "MOVIE", AgeCertification: "PG-13"},
		{ID: "2", Title: "Movie 2", Type: "MOVIE", AgeCertification: "R"},
		{ID: "3", Title: "Show 1", Type: "SHOW", AgeCertification: "PG-13"},
		{ID: "4", Title: "Movie 3", Type: "MOVIE", AgeCertification: "PG"},
		{ID: "5", Title: "Movie 4", Type: "MOVIE", AgeCertification: "R"},
		{ID: "6", Title: "Show 2", Type: "SHOW", AgeCertification: "PG"},
	}

	expected := map[string]int{
		"PG-13": 1,
		"R":     2,
		"PG":    1,
	}

	result := ListMoviesCountByAgeCertificate(titles)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ListMoviesCountByAgeCertificate() returned unexpected result:\nExpected: %v\nActual: %v", expected, result)
	}
}

func TestListMovieCountByRuntime(t *testing.T) {
	titles := []Title{
		{ID: "1", Title: "Movie 1", Type: "MOVIE", Runtime: 90},
		{ID: "2", Title: "Movie 2", Type: "MOVIE", Runtime: 120},
		{ID: "3", Title: "Show 1", Type: "SHOW", Runtime: 45},
		{ID: "4", Title: "Movie 3", Type: "MOVIE", Runtime: 180},
		{ID: "5", Title: "Movie 4", Type: "MOVIE", Runtime: 300},
		{ID: "6", Title: "Show 2", Type: "MOVIE", Runtime: 20},
	}

	expected := map[int]int{
		30:   1,
		45:   1,
		60:   1,
		120:  2,
		360:  5,
		1000: 5,
	}

	result := ListMovieCountByRuntime(titles)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ListMovieCountByRuntime() returned unexpected result:\nExpected: %v\nActual: %v", expected, result)
	}
}

func TestListTitlesCountPercentageByGenres(t *testing.T) {
	titles := []Title{
		{ID: "1", Title: "Movie 1", Genres: []string{"['action', 'adventure']"}},
		{ID: "2", Title: "Movie 2", Genres: []string{"['action', 'drama']"}},
		{ID: "3", Title: "Movie 3", Genres: []string{"['comedy', 'drama']"}},
		{ID: "4", Title: "Show 1", Genres: []string{"['drama']"}},
		{ID: "5", Title: "Movie 4", Genres: []string{"['comedy', 'romance']"}},
		{ID: "6", Title: "Show 2", Genres: []string{"['action', 'comedy']"}},
	}

	expectedGenreCount := map[string]int{
		"action":    3,
		"adventure": 1,
		"drama":     3,
		"comedy":    3,
		"romance":   1,
	}

	resultGenreCount, _ := ListTitlesCountPercentageByGenres(titles)

	if !reflect.DeepEqual(resultGenreCount, expectedGenreCount) {
		t.Errorf("ListTitlesCountPercentageByGenres() returned unexpected genre count:\nExpected: %v\nActual: %v", expectedGenreCount, resultGenreCount)
	}
}

func TestListTitlesCountPercentageByCountry(t *testing.T) {
	titles := []Title{
		{ID: "1", Title: "Movie 1", ProductionCountries: []string{"['US', 'UK']"}},
		{ID: "2", Title: "Movie 2", ProductionCountries: []string{"['US', 'Canada']"}},
		{ID: "3", Title: "Movie 3", ProductionCountries: []string{"['UK']"}},
		{ID: "4", Title: "Show 1", ProductionCountries: []string{"['Canada']"}},
		{ID: "5", Title: "Movie 4", ProductionCountries: []string{"['US']"}},
		{ID: "6", Title: "Show 2", ProductionCountries: []string{"['UK', 'Canada']"}},
	}

	expectedCountryCount := map[string]int{
		"US":     3,
		"UK":     3,
		"Canada": 3,
	}

	resultCountryCount, _ := ListTitlesCountPercentageByCountry(titles)

	if !reflect.DeepEqual(resultCountryCount, expectedCountryCount) {
		t.Errorf("ListTitlesCountPercentageByCountry() returned unexpected country count:\nExpected: %v\nActual: %v", expectedCountryCount, resultCountryCount)
	}
}

func TestListTitleCountBySeasons(t *testing.T) {
	titles := []Title{
		{ID: "1", Title: "Show 1", Seasons: 3},
		{ID: "2", Title: "Show 2", Seasons: 2},
		{ID: "3", Title: "Movie 1", Seasons: 0},
		{ID: "4", Title: "Show 3", Seasons: 1},
		{ID: "5", Title: "Movie 2", Seasons: 0},
	}

	expectedSeasonsCounts := map[int]int{
		3: 1,
		2: 1,
		1: 1,
	}

	resultSeasonsCounts := ListTitleCountBySeasons(titles)

	if !reflect.DeepEqual(resultSeasonsCounts, expectedSeasonsCounts) {
		t.Errorf("ListTitleCountBySeasons() returned unexpected seasons counts:\nExpected: %v\nActual: %v", expectedSeasonsCounts, resultSeasonsCounts)
	}
}

func TestListTitlesCountByIMDbScore(t *testing.T) {
	titles := []Title{
		{ID: "1", Title: "Movie 1", IMDbScore: 8.2},
		{ID: "2", Title: "Movie 2", IMDbScore: 7.5},
		{ID: "3", Title: "Movie 3", IMDbScore: 6.8},
		{ID: "4", Title: "Movie 4", IMDbScore: 4.3},
		{ID: "5", Title: "Movie 5", IMDbScore: 9.1},
		{ID: "6", Title: "Movie 6", IMDbScore: 5.5},
	}

	expectedTitlesCountByImdbMap := map[[2]int]int{
		{0, 4}:  0,
		{4, 6}:  2,
		{6, 8}:  2,
		{8, 10}: 2,
	}

	resultTitlesCountByImdbMap := ListTitlesCountByIMDbScore(titles)

	if !reflect.DeepEqual(resultTitlesCountByImdbMap, expectedTitlesCountByImdbMap) {
		t.Errorf("ListTitlesCountByIMDbScore() returned unexpected titles count by IMDb score:\nExpected: %v\nActual: %v", expectedTitlesCountByImdbMap, resultTitlesCountByImdbMap)
	}
}

func TestListTitlesCountByRuntime(t *testing.T) {
	titles := []Title{
		{ID: "1", Title: "Movie 1", Runtime: 90},
		{ID: "2", Title: "Movie 2", Runtime: 120},
		{ID: "3", Title: "Movie 3", Runtime: 75},
		{ID: "4", Title: "Movie 4", Runtime: 150},
		{ID: "5", Title: "Movie 5", Runtime: 180},
		{ID: "6", Title: "Movie 6", Runtime: 45},
	}

	expectedTitlesCountByRuntimeMap := map[int]int{
		0:   6,
		30:  6,
		45:  5,
		60:  5,
		120: 2,
		360: 0,
	}

	resultTitlesCountByRuntimeMap := ListTitlesCountByRuntime(titles)

	if !reflect.DeepEqual(resultTitlesCountByRuntimeMap, expectedTitlesCountByRuntimeMap) {
		t.Errorf("ListTitlesCountByRuntime() returned unexpected titles count by runtime:\nExpected: %v\nActual: %v", expectedTitlesCountByRuntimeMap, resultTitlesCountByRuntimeMap)
	}
}

func TestGetTitleTypeCountsAndPercentages(t *testing.T) {
	titles := []Title{
		{ID: "1", Type: "MOVIE"},
		{ID: "2", Type: "MOVIE"},
		{ID: "3", Type: "SHOW"},
		{ID: "4", Type: "MOVIE"},
		{ID: "5", Type: "SHOW"},
	}

	expectedMovieCount := 3
	expectedMoviePercentage := float64(60)
	expectedShowCount := 2
	expectedShowPercentage := float64(40)

	resultMovieCount, resultMoviePercentage, resultShowCount, resultShowPercentage := GetTitleTypeCountsAndPercentages(titles)

	// Check movie count
	if resultMovieCount != expectedMovieCount {
		t.Errorf("GetTitleTypeCountsAndPercentages() returned unexpected movie count:\nExpected: %v\nActual: %v", expectedMovieCount, resultMovieCount)
	}

	// Check movie percentage
	if math.Abs(resultMoviePercentage-expectedMoviePercentage) > 0.001 {
		t.Errorf("GetTitleTypeCountsAndPercentages() returned unexpected movie percentage:\nExpected: %v\nActual: %v", expectedMoviePercentage, resultMoviePercentage)
	}

	// Check show count
	if resultShowCount != expectedShowCount {
		t.Errorf("GetTitleTypeCountsAndPercentages() returned unexpected show count:\nExpected: %v\nActual: %v", expectedShowCount, resultShowCount)
	}

	// Check show percentage
	if math.Abs(resultShowPercentage-expectedShowPercentage) > 0.001 {
		t.Errorf("GetTitleTypeCountsAndPercentages() returned unexpected show percentage:\nExpected: %v\nActual: %v", expectedShowPercentage, resultShowPercentage)
	}
}
