package middlewares

import (
	"github.com/doug-martin/goqu/v9"
	"go.uber.org/zap"

	"github.com/Improwised/golang-api/config"
	"github.com/Improwised/golang-api/models"
)

type Middleware struct {
	config config.AppConfig
	logger *zap.Logger
}

type TitleCreditCheckMiddleware struct {
	titleModel  *models.TitleModel
	creditModel *models.CreditModel
	logger      *zap.Logger
}

func NewTitleCreditCheckMiddleware(goqu *goqu.Database, logger *zap.Logger) (*TitleCreditCheckMiddleware, error) {
	titleModel, err := models.InitTitleModel(goqu)
	if err != nil {
		return nil, err
	}
	creditModel, err := models.InitCreditModel(goqu)
	if err != nil {
		return nil, err
	}

	return &TitleCreditCheckMiddleware{
		titleModel:  titleModel,
		creditModel: creditModel,
		logger:      logger,
	}, nil
}

func NewMiddleware(cfg config.AppConfig, logger *zap.Logger) Middleware {
	return Middleware{
		config: cfg,
		logger: logger,
	}
}
