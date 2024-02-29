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
