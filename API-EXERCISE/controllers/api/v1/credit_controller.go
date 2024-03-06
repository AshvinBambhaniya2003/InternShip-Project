package v1

import (
	"database/sql"
	"encoding/json"
	"net/http"

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

// CreditController for credit controllers
type CreditController struct {
	titleModel  *models.TitleModel
	creditModel *models.CreditModel
	logger      *zap.Logger
	event       *events.Events
	pub         *watermill.WatermillPublisher
}

func NewCreditController(goqu *goqu.Database, logger *zap.Logger, event *events.Events, pub *watermill.WatermillPublisher) (*CreditController, error) {
	creditModel, err := models.InitCreditModel(goqu)
	if err != nil {
		return nil, err
	}

	titleModel, err := models.InitTitleModel(goqu)
	if err != nil {
		return nil, err
	}

	return &CreditController{
		titleModel:  titleModel,
		creditModel: creditModel,
		logger:      logger,
		event:       event,
		pub:         pub,
	}, nil
}

// CreateCredit create a credit
// swagger:route POST /titles/{titleId}/credits Credits RequestCreateCredit
//
// Create a credit.
//
//		Consumes:
//		- application/json
//
//		Schemes: http
//
//		Responses:
//		  201: ResponseCredit
//	   400: GenericResFailBadRequest
//		  500: GenericResError
//	   404: GenericResFailNotFound
func (ctrl *CreditController) Create(c *fiber.Ctx) error {

	titleID := c.Params(constants.ParamTitleId)

	var creditReq structs.ReqRegisterCredit

	err := json.Unmarshal(c.Body(), &creditReq)
	if err != nil {
		return utils.JSONFail(c, http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(creditReq)
	if err != nil {
		return utils.JSONFail(c, http.StatusBadRequest, utils.ValidatorErrorString(err))
	}

	id := xid.New().String()

	credit, err := ctrl.creditModel.Insert(models.Credit{
		Id:        id,
		PersonID:  creditReq.PersonID,
		TitleID:   titleID,
		Name:      creditReq.Name,
		Character: creditReq.Character,
		Role:      creditReq.Role,
	})

	if err != nil {
		ctrl.logger.Error("error while insert credit", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, "error while creating Credit, please try after sometime")
	}

	return utils.JSONSuccess(c, http.StatusCreated, credit)
}

// ListCredit list credit by titleId
// swagger:route GET /titles/{titleId}/credits Credits RequestGetCreditTitle
//
// List credits by titleId.
//
//		Consumes:
//		- application/json
//
//		Schemes: http
//
//		Responses:
//		  200: ResponseListCredit
//	   404: GenericResFailNotFound
//		  500: GenericResError
func (ctrl *CreditController) ListByTitleId(c *fiber.Ctx) error {

	titleId := c.Params(constants.ParamTitleId)

	credits, err := ctrl.creditModel.ListCredits(titleId)
	if err != nil {
		ctrl.logger.Error("error while getting credits", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, "Error while get credit list")
	}
	return utils.JSONSuccess(c, http.StatusOK, credits)
}

// GetCredit get a credit by id
// swagger:route GET /titles/{titleId}/credits/{creditId} Credits RequestGetCredit
//
// Get a credit.
//
//		Consumes:
//		- application/json
//
//		Schemes: http
//
//		Responses:
//		  200: ResponseCredit
//	   404: GenericResFailNotFound
//		  500: GenericResError
func (ctrl *CreditController) GetById(c *fiber.Ctx) error {

	id := c.Params(constants.ParamCreditId)

	credit, err := ctrl.creditModel.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.JSONFail(c, http.StatusNotFound, "Credit Does not exist")
		}
		ctrl.logger.Error("error while get credit by id", zap.Any("id", id), zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, "Error while Get Credit")
	}

	return utils.JSONSuccess(c, http.StatusOK, credit)
}

// DeleteCredit delete a credit by id
// swagger:route DELETE /titles/{titleId}/credits/{creditId} Credits RequestDeleteCredit
//
// Delete a credit.
//
//		Consumes:
//		- application/json
//
//		Schemes: http
//
//		Responses:
//		  200: ResponseDeleteCredit
//	   404: GenericResFailNotFound
//		  500: GenericResError
func (ctrl *CreditController) Delete(c *fiber.Ctx) error {

	id := c.Params(constants.ParamCreditId)

	err := ctrl.creditModel.Delete(id)
	if err != nil {
		ctrl.logger.Error("error while delete credit by id", zap.Any("id", id), zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, "Error while Delete title")
	}

	return utils.JSONSuccess(c, http.StatusOK, "Suceesfully Deleted")
}
