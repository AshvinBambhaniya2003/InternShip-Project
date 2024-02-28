package v1

import (
	"net/http"

	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/Improwised/golang-api/constants"
	"github.com/Improwised/golang-api/models"
	"github.com/Improwised/golang-api/pkg/events"
	"github.com/Improwised/golang-api/pkg/watermill"
	"github.com/Improwised/golang-api/utils"
)

// TitleController for title controllers
type TitleController struct {
	titleModel *models.TitleModel
	logger     *zap.Logger
	event      *events.Events
	pub        *watermill.WatermillPublisher
}

func NewTitleController(goqu *goqu.Database, logger *zap.Logger, event *events.Events, pub *watermill.WatermillPublisher) (*TitleController, error) {
	titleModel, err := models.InitTitleModel(goqu)
	if err != nil {
		return nil, err
	}

	return &TitleController{
		titleModel: titleModel,
		logger:     logger,
		event:      event,
		pub:        pub,
	}, nil
}

func (ctrl *TitleController) List(c *fiber.Ctx) error {
	titles, err := ctrl.titleModel.List()
	if err != nil {
		return utils.JSONError(c, http.StatusInternalServerError, "Error while get title list")
	}
	return utils.JSONSuccess(c, http.StatusOK, titles)
}

func (ctrl *TitleController) GetById(c *fiber.Ctx) error {

	titleID := c.Params(constants.ParamTitleId)

	title, err := ctrl.titleModel.GetById(titleID)
	if err != nil {
		return utils.JSONError(c, http.StatusInternalServerError, "no any title associate with given id")
	}

	return utils.JSONSuccess(c, http.StatusOK, title)
}

func (ctrl *TitleController) Delete(c *fiber.Ctx) error {

	titleID := c.Params(constants.ParamTitleId)

	title, err := ctrl.titleModel.GetById(titleID)
	if err != nil {
		return utils.JSONError(c, http.StatusInternalServerError, "no any title associate with given id")
	}

	err = ctrl.titleModel.DeleteTitle(titleID)
	if err != nil {
		return utils.JSONError(c, http.StatusInternalServerError, "Error while Delete title")
	}
	return utils.JSONSuccess(c, http.StatusOK, title)
}
