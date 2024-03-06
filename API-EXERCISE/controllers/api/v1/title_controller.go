package v1

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
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

// TitleList list the titles
// swagger:route GET /titles Titles listTitle
//
//	List Titles
//
//	List Titles
//
//		Consumes:
//		- application/json
//
//		Schemes: http
//
//		Responses:
//		  200: ResponseTitleList
//	   400: GenericResFailBadRequest
//		  500: GenericResError
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

// TitleGet get the title by id
// swagger:route GET /titles/{titleId} Titles RequestGetTitle
//
// Get a title.
//
//		Consumes:
//		- application/json
//
//		Schemes: http
//
//		Responses:
//		  200: ResponseTitle
//	   404: GenericResFailNotFound
//		  500: GenericResError
func (ctrl *TitleController) GetById(c *fiber.Ctx) error {

	titleID := c.Params(constants.ParamTitleId)

	title, err := ctrl.titleModel.GetById(titleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.JSONFail(c, http.StatusNotFound, constants.TitleNotExist)
		}
		ctrl.logger.Error("error while get title by id", zap.Any("id", titleID), zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, constants.ErrGetTitle)
	}

	return utils.JSONSuccess(c, http.StatusOK, title)
}

// TitleDelete delete the title by id
// swagger:route DELETE /titles/{titleId} Titles RequestDeleteTitle
//
// Delete a title.
//
//		Consumes:
//		- application/json
//
//		Schemes: http
//
//		Responses:
//		  200: ResponseTitle
//	   404: GenericResFailNotFound
//		  500: GenericResError
func (ctrl *TitleController) Delete(c *fiber.Ctx) error {

	titleID := c.Params(constants.ParamTitleId)

	title, err := ctrl.titleModel.GetById(titleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.JSONFail(c, http.StatusNotFound, constants.TitleNotExist)
		}
		ctrl.logger.Error("error while get title by id", zap.Any("id", titleID), zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, constants.ErrGetTitle)
	}

	err = ctrl.titleModel.Delete(titleID)
	if err != nil {
		ctrl.logger.Error("error while delete title", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, "error while delete title")
	}
	return utils.JSONSuccess(c, http.StatusOK, title)
}

// CreateTitle create a title
// swagger:route POST /titles Titles RequestCreateTitle
//
// Create a title.
//
//		Consumes:
//		- application/json
//
//		Schemes: http
//
//		Responses:
//		  201: ResponseTitle
//	   400: GenericResFailBadRequest
//		  500: GenericResError
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

	id := xid.New().String()

	title, err := ctrl.titleModel.Create(models.Title{
		ID:                id,
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

// UpdateTitle Update a title
// swagger:route PUT /titles/{titleId} Titles RequestUpdateTitle
//
// Update a title.
//
//		Consumes:
//		- application/json
//
//		Schemes: http
//
//		Responses:
//		  200: ResponseTitle
//	   400: GenericResFailBadRequest
//		  500: GenericResError
//		404: GenericResFailNotFound
func (ctrl *TitleController) Update(c *fiber.Ctx) error {

	titleID := c.Params(constants.ParamTitleId)

	_, err := ctrl.titleModel.GetById(titleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.JSONFail(c, http.StatusNotFound, constants.TitleNotExist)
		}
		ctrl.logger.Error("error while get title by id", zap.Any("id", titleID), zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, constants.ErrGetTitle)
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

	title, err := ctrl.titleModel.Update(titleID, models.Title{
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
	})
	if err != nil {
		ctrl.logger.Error("error while update title by id", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, "Error while update title")
	}
	return utils.JSONSuccess(c, http.StatusOK, title)
}
