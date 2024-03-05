package v1_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Improwised/golang-api/models"
	"github.com/Improwised/golang-api/pkg/structs"
)

type TitleResponse struct {
	Status string       `json:"status"`
	Data   models.Title `json:"data"`
}

type TitleListResponse struct {
	Status string         `json:"status"`
	Data   []models.Title `json:"data"`
}

var titleReq = structs.ReqRegisterTitle{
	Title:             "12th fail",
	Type:              "movie",
	Description:       "This is an example description of the title.",
	ReleaseYear:       2022,
	AgeCertification:  "PG-13",
	Runtime:           120,
	Genres:            "['Action', 'Adventure', 'Sci-Fi']",
	ProductionCountry: "['United States', 'United Kingdom']",
	Seasons:           45,
	IMDBID:            "tt1234567",
	IMDBScore:         7.5,
	IMDBVotes:         10000,
	TMDBPopularity:    8.5,
	TMDBScore:         7.8,
}

func compareTitles(expected structs.ReqRegisterTitle, actual models.Title) bool {
	// Compare each field except CreatedAt and UpdatedAt
	return expected.Title == actual.Title &&
		expected.Type == actual.Type &&
		expected.Description == actual.Description &&
		expected.ReleaseYear == actual.ReleaseYear &&
		expected.AgeCertification == actual.AgeCertification &&
		expected.Runtime == actual.Runtime &&
		expected.Genres == actual.Genres &&
		expected.ProductionCountry == actual.ProductionCountry &&
		expected.Seasons == actual.Seasons &&
		expected.IMDBID == actual.IMDBID &&
		expected.IMDBScore == actual.IMDBScore &&
		expected.IMDBVotes == actual.IMDBVotes &&
		expected.TMDBPopularity == actual.TMDBPopularity &&
		expected.TMDBScore == actual.TMDBScore
}

func TestCreateTitle(t *testing.T) {
	t.Run("create title with required fields are not provided", func(t *testing.T) {
		req := structs.ReqRegisterTitle{
			Title: "friends",
			Type:  "Show",
		}
		res, err := client.
			R().
			EnableTrace().
			SetBody(req).
			Post("/api/v1/titles")

		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode())
	})

	t.Run("create title with valid input", func(t *testing.T) {
		res, err := client.
			R().
			EnableTrace().
			SetBody(titleReq).
			Post("/api/v1/titles")
		assert.Nil(t, err)

		var titleResponse TitleResponse
		err = json.Unmarshal(res.Body(), &titleResponse)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusCreated, res.StatusCode())
		assert.True(t, compareTitles(titleReq, titleResponse.Data))

	})

	t.Cleanup(func() {
		_, err := db.Exec("delete from titles where title='12th fail'")
		assert.Nil(t, err)
	})

}

func TestListTitle(t *testing.T) {

	t.Run("create title with valid input", func(t *testing.T) {
		for i := 0; i < 4; i++ {
			res, err := client.
				R().
				EnableTrace().
				SetBody(titleReq).
				Post("/api/v1/titles")

			assert.Nil(t, err)
			assert.Equal(t, http.StatusCreated, res.StatusCode())
		}

	})

	t.Run("list titles", func(t *testing.T) {
		res, err := client.
			R().
			EnableTrace().
			Get("/api/v1/titles")

		assert.Nil(t, err)

		var titleListResponse TitleListResponse
		err = json.Unmarshal(res.Body(), &titleListResponse)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode())

		expectedCount := 4
		assert.Len(t, titleListResponse.Data, expectedCount)
	})

	t.Cleanup(func() {
		_, err := db.Exec("delete from titles where title='12th fail'")
		assert.Nil(t, err)
	})

}

func TestGetTitle(t *testing.T) {
	var id string

	t.Run("create title with valid input", func(t *testing.T) {
		res, err := client.
			R().
			EnableTrace().
			SetBody(titleReq).
			Post("/api/v1/titles")

		assert.Nil(t, err)

		var responseData TitleResponse
		err = json.Unmarshal(res.Body(), &responseData)
		assert.Nil(t, err)
		id = responseData.Data.ID

		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode())
	})

	t.Run("get title", func(t *testing.T) {
		res, err := client.
			R().
			EnableTrace().
			Get(fmt.Sprintf("/api/v1/titles/%s", id))
		assert.Nil(t, err)

		var responseData TitleResponse
		err = json.Unmarshal(res.Body(), &responseData)
		assert.Nil(t, err)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode())
		assert.True(t, compareTitles(titleReq, responseData.Data))
	})

	t.Cleanup(func() {
		_, err := db.Exec(fmt.Sprintf("delete from titles where id='%s'", id))
		assert.Nil(t, err)
	})

}

func TestDeleteTitle(t *testing.T) {
	var id string

	t.Run("create title with valid input", func(t *testing.T) {
		res, err := client.
			R().
			EnableTrace().
			SetBody(titleReq).
			Post("/api/v1/titles")

		assert.Nil(t, err)

		var responseData TitleResponse
		err = json.Unmarshal(res.Body(), &responseData)
		assert.Nil(t, err)
		id = responseData.Data.ID

		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode())
	})

	t.Run("delete title", func(t *testing.T) {
		res, err := client.
			R().
			EnableTrace().
			Delete(fmt.Sprintf("/api/v1/titles/%s", id))

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode())

		res, err = client.
			R().
			EnableTrace().
			Get("/api/v1/titles")

		assert.Nil(t, err)

		var titleListResponse TitleListResponse
		err = json.Unmarshal(res.Body(), &titleListResponse)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode())

		expectedCount := 0
		assert.Len(t, titleListResponse.Data, expectedCount)
	})

}

func TestUpdateTitle(t *testing.T) {
	var id string

	t.Run("create title with valid input", func(t *testing.T) {
		res, err := client.
			R().
			EnableTrace().
			SetBody(titleReq).
			Post("/api/v1/titles")

		assert.Nil(t, err)

		var responseData TitleResponse
		err = json.Unmarshal(res.Body(), &responseData)
		assert.Nil(t, err)

		id = responseData.Data.ID

		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode())
	})

	t.Run("update title", func(t *testing.T) {
		req := structs.ReqRegisterTitle{
			Title:             "12th Fail",
			Type:              "movie",
			Description:       "This is an example description of the title.",
			ReleaseYear:       2024,
			AgeCertification:  "PG-13",
			Runtime:           146,
			Genres:            "['Biographi', 'Drama']",
			ProductionCountry: "['India']",
			Seasons:           45,
			IMDBID:            "tt1234567",
			IMDBScore:         9.1,
			IMDBVotes:         10000,
			TMDBPopularity:    8.5,
			TMDBScore:         7.8,
		}
		res, err := client.
			R().
			EnableTrace().
			SetBody(req).
			Put(fmt.Sprintf("/api/v1/titles/%s", id))

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode())

		res, err = client.
			R().
			EnableTrace().
			Get(fmt.Sprintf("/api/v1/titles/%s", id))
		assert.Nil(t, err)

		var responseData TitleResponse
		err = json.Unmarshal(res.Body(), &responseData)
		assert.Nil(t, err)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode())
		assert.True(t, compareTitles(req, responseData.Data))
	})

	t.Cleanup(func() {
		_, err := db.Exec(fmt.Sprintf("delete from titles where id='%s'", id))
		assert.Nil(t, err)
	})

}
