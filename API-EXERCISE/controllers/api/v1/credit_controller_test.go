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

type CreditResponse struct {
	Status string        `json:"status"`
	Data   models.Credit `json:"data"`
}

type CreditListResponse struct {
	Status string          `json:"status"`
	Data   []models.Credit `json:"data"`
}

var creditReq = structs.ReqRegisterCredit{
	PersonID:  45,
	Name:      "Ashvin",
	Character: "some char",
	Role:      "Actor",
}

func compareCredit(expected structs.ReqRegisterCredit, actual models.Credit) bool {
	// Compare each field except CreatedAt and UpdatedAt
	return expected.PersonID == actual.PersonID &&
		expected.Name == actual.Name &&
		expected.Character == actual.Character &&
		expected.Role == actual.Role
}

func TestCreateCredit(t *testing.T) {
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
	t.Run("create credit with required fields are not provided", func(t *testing.T) {
		req := structs.ReqRegisterCredit{
			Name: "someonename",
			Role: "somerole",
		}
		res, err := client.
			R().
			EnableTrace().
			SetBody(req).
			Post(fmt.Sprintf("api/v1/titles/%s/credits", id))

		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode())
	})

	t.Run("create credit with valid input", func(t *testing.T) {
		res, err := client.
			R().
			EnableTrace().
			SetBody(creditReq).
			Post(fmt.Sprintf("api/v1/titles/%s/credits", id))
		assert.Nil(t, err)

		var creditResponse CreditResponse
		err = json.Unmarshal(res.Body(), &creditResponse)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusCreated, res.StatusCode())
		assert.True(t, compareCredit(creditReq, creditResponse.Data))

	})

	t.Cleanup(func() {
		_, err := db.Exec("delete from credits where name='Ashvin'")
		assert.Nil(t, err)
		_, err = db.Exec("delete from titles where title='12th fail'")
		assert.Nil(t, err)
	})

}

func TestListCredit(t *testing.T) {

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

	t.Run("create credits with valid input", func(t *testing.T) {
		for i := 0; i < 4; i++ {
			res, err := client.
				R().
				EnableTrace().
				SetBody(creditReq).
				Post(fmt.Sprintf("api/v1/titles/%s/credits", id))

			assert.Nil(t, err)
			assert.Equal(t, http.StatusCreated, res.StatusCode())
		}

	})

	t.Run("list credits", func(t *testing.T) {
		res, err := client.
			R().
			EnableTrace().
			Get(fmt.Sprintf("api/v1/titles/%s/credits", id))

		assert.Nil(t, err)

		var creditListResponse CreditListResponse
		err = json.Unmarshal(res.Body(), &creditListResponse)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode())

		expectedCount := 4
		assert.Len(t, creditListResponse.Data, expectedCount)
	})

	t.Cleanup(func() {
		_, err := db.Exec("delete from credits where name='Ashvin'")
		assert.Nil(t, err)
		_, err = db.Exec("delete from titles where title='12th fail'")
		assert.Nil(t, err)
	})

}

func TestGetCredit(t *testing.T) {

	var titleId string
	var creditId string

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
		titleId = responseData.Data.ID

		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode())

	})

	t.Run("create credit with valid input", func(t *testing.T) {
		res, err := client.
			R().
			EnableTrace().
			SetBody(creditReq).
			Post(fmt.Sprintf("api/v1/titles/%s/credits", titleId))
		assert.Nil(t, err)

		var creditResponse CreditResponse
		err = json.Unmarshal(res.Body(), &creditResponse)
		assert.Nil(t, err)
		creditId = creditResponse.Data.Id

		assert.Equal(t, http.StatusCreated, res.StatusCode())
		assert.True(t, compareCredit(creditReq, creditResponse.Data))

	})

	t.Run("get credit", func(t *testing.T) {
		res, err := client.
			R().
			EnableTrace().
			Get(fmt.Sprintf("api/v1/titles/%s/credits/%s", titleId, creditId))
		assert.Nil(t, err)

		var responseData CreditResponse
		err = json.Unmarshal(res.Body(), &responseData)
		assert.Nil(t, err)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode())
		assert.True(t, compareCredit(creditReq, responseData.Data))
	})

	t.Cleanup(func() {
		_, err := db.Exec("delete from credits where name='Ashvin'")
		assert.Nil(t, err)
		_, err = db.Exec("delete from titles where title='12th fail'")
		assert.Nil(t, err)
	})

}

func TestDeleteCredit(t *testing.T) {

	var titleId string
	var creditId string

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
		titleId = responseData.Data.ID

		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode())

	})

	t.Run("create credit with valid input", func(t *testing.T) {
		res, err := client.
			R().
			EnableTrace().
			SetBody(creditReq).
			Post(fmt.Sprintf("api/v1/titles/%s/credits", titleId))
		assert.Nil(t, err)

		var creditResponse CreditResponse
		err = json.Unmarshal(res.Body(), &creditResponse)
		assert.Nil(t, err)
		creditId = creditResponse.Data.Id

		assert.Equal(t, http.StatusCreated, res.StatusCode())
		assert.True(t, compareCredit(creditReq, creditResponse.Data))

	})

	t.Run("delete credit", func(t *testing.T) {
		res, err := client.
			R().
			EnableTrace().
			Delete(fmt.Sprintf("api/v1/titles/%s/credits/%s", titleId, creditId))

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode())
	})

	t.Cleanup(func() {
		_, err := db.Exec("delete from credits where name='Ashvin'")
		assert.Nil(t, err)
		_, err = db.Exec("delete from titles where title='12th fail'")
		assert.Nil(t, err)
	})

}
