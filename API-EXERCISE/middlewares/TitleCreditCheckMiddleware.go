package middlewares

import (
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/Improwised/golang-api/constants"
	"github.com/Improwised/golang-api/utils"
)

func (t *TitleCreditCheckMiddleware) TitleExist(c *fiber.Ctx) error {

	id := c.Params(constants.ParamTitleId)

	_, err := t.titleModel.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.JSONFail(c, http.StatusNotFound, "Title Does not exist")
		}
		return utils.JSONError(c, http.StatusInternalServerError, "Error while Get Title")
	}

	return c.Next()
}

func (t *TitleCreditCheckMiddleware) CreditTitleVerification(c *fiber.Ctx) error {

	titleId := c.Params(constants.ParamTitleId)

	title, err := t.titleModel.GetById(titleId)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.JSONFail(c, http.StatusNotFound, "Title Does not exist")
		}
		return utils.JSONError(c, http.StatusInternalServerError, "Error while Get Title")
	}

	creditId := c.Params(constants.ParamCreditId)

	credit, err := t.creditModel.GetById(creditId)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.JSONFail(c, http.StatusNotFound, "Credit Does not exist")
		}
		return utils.JSONError(c, http.StatusInternalServerError, "error while get Credit")
	}

	if title.ID != credit.TitleID {
		return utils.JSONError(c, http.StatusBadRequest, "Your titleId is not assosiate with given creditId")
	}

	return c.Next()
}
