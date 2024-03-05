package v1

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"

	"github.com/Improwised/golang-api/constants"
	"github.com/Improwised/golang-api/models"
	"github.com/Improwised/golang-api/pkg/events"
	"github.com/Improwised/golang-api/pkg/structs"
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

	queries := c.Queries()
	pageStr := c.Query(constants.Page)

	var pageInt int
	var err error
	if pageStr != "" {
		pageInt, err = strconv.Atoi(pageStr)
		if pageInt <= 0 {
			pageInt = 1
		}

		if err != nil {
			return utils.JSONError(c, http.StatusBadRequest, "please enter page in int")
		}
	}

	titles, err := ctrl.titleModel.List(queries, pageInt)
	if err != nil {
		ctrl.logger.Error("error while getting titles", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, "error while getting titles")
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

	err = ctrl.titleModel.Delete(titleID)
	if err != nil {
		return utils.JSONError(c, http.StatusInternalServerError, "Error while Delete title")
	}
	return utils.JSONSuccess(c, http.StatusOK, title)
}

func (ctrl *TitleController) Create(c *fiber.Ctx) error {

	var titleReq structs.ReqRegisterTitle

	err := json.Unmarshal(c.Body(), &titleReq)
	if err != nil {
		return utils.JSONFail(c, http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(titleReq)
	if err != nil {
		return utils.JSONFail(c, http.StatusBadRequest, utils.ValidatorErrorString(err))
	}

	title, err := ctrl.titleModel.Insert(models.Title{
		Title:             titleReq.Title,
		Type:              titleReq.Type,
		Description:       titleReq.Description,
		ReleaseYear:       titleReq.ReleaseYear,
		AgeCertification:  titleReq.AgeCertification,
		Runtime:           titleReq.Runtime,
		Genres:            titleReq.Genres,
		ProductionCountry: titleReq.ProductionCountry,
		Seasons:           titleReq.Seasons,
		IMDBID:            titleReq.IMDBID,
		IMDBScore:         titleReq.IMDBScore,
		IMDBVotes:         titleReq.IMDBVotes,
		TMDBPopularity:    titleReq.TMDBPopularity,
		TMDBScore:         titleReq.TMDBScore,
	})

	if err != nil {
		ctrl.logger.Error("error while insert title", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, "error while creating title, please try after sometime")
	}

	return utils.JSONSuccess(c, http.StatusCreated, title)
}

func (ctrl *TitleController) Update(c *fiber.Ctx) error {
	titleID := c.Params(constants.ParamTitleId)

	title, err := ctrl.titleModel.GetById(titleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.JSONFail(c, http.StatusNotFound, "title not exist")
		}
		return utils.JSONError(c, http.StatusInternalServerError, "no any title associate with given id")
	}

	var titleReq structs.ReqRegisterTitle

	err = json.Unmarshal(c.Body(), &titleReq)
	if err != nil {
		return utils.JSONFail(c, http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(titleReq)
	if err != nil {
		return utils.JSONFail(c, http.StatusBadRequest, utils.ValidatorErrorString(err))
	}

	now := time.Now()
	timestampString := now.Format("2006-01-02T15:04:05.999999Z")
	err = ctrl.titleModel.Update(titleID, models.Title{
		ID:                titleID,
		Title:             titleReq.Title,
		Type:              titleReq.Type,
		Description:       titleReq.Description,
		ReleaseYear:       titleReq.ReleaseYear,
		AgeCertification:  titleReq.AgeCertification,
		Runtime:           titleReq.Runtime,
		Genres:            titleReq.Genres,
		ProductionCountry: titleReq.ProductionCountry,
		Seasons:           titleReq.Seasons,
		IMDBID:            titleReq.IMDBID,
		IMDBScore:         titleReq.IMDBScore,
		IMDBVotes:         titleReq.IMDBVotes,
		TMDBPopularity:    titleReq.TMDBPopularity,
		TMDBScore:         titleReq.TMDBScore,
		CreatedAt:         title.CreatedAt,
		UpdatedAt:         timestampString,
	})
	if err != nil {
		return utils.JSONError(c, http.StatusInternalServerError, "Error while Update title")
	}
	return utils.JSONSuccess(c, http.StatusOK, titleReq)
}
