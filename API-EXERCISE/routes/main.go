package routes

import (
	"fmt"
	"sync"

	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/Improwised/golang-api/config"
	"github.com/Improwised/golang-api/constants"
	controller "github.com/Improwised/golang-api/controllers/api/v1"
	"github.com/Improwised/golang-api/middlewares"
	"github.com/Improwised/golang-api/pkg/events"
	pMetrics "github.com/Improwised/golang-api/pkg/prometheus"
	"github.com/Improwised/golang-api/pkg/watermill"
)

var mu sync.Mutex

// Setup func
func Setup(app *fiber.App, goqu *goqu.Database, logger *zap.Logger, config config.AppConfig, events *events.Events, pMetrics *pMetrics.PrometheusMetrics, pub *watermill.WatermillPublisher) error {
	mu.Lock()

	app.Use(middlewares.LogHandler(logger, pMetrics))

	app.Static("/assets/", "./assets")

	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Render("./assets/index.html", fiber.Map{})
	})

	router := app.Group("/api")
	v1 := router.Group("/v1")

	titleCreditCheckMiddlewares, err := middlewares.NewTitleCreditCheckMiddleware(goqu, logger)
	if err != nil {
		return err
	}

	middlewares := middlewares.NewMiddleware(config, logger)

	err = setupAuthController(v1, goqu, logger, middlewares, config)
	if err != nil {
		return err
	}

	err = setupUserController(v1, goqu, logger, middlewares, events, pub)
	if err != nil {
		return err
	}

	err = setupTitleController(v1, goqu, logger, middlewares, events, pub)
	if err != nil {
		return err
	}

	err = setupCreditController(v1, goqu, logger, titleCreditCheckMiddlewares, events, pub)
	if err != nil {
		return err
	}

	err = healthCheckController(app, goqu, logger)
	if err != nil {
		return err
	}

	err = metricsController(app, goqu, logger, pMetrics)
	if err != nil {
		return err
	}

	mu.Unlock()
	return nil
}

func setupAuthController(v1 fiber.Router, goqu *goqu.Database, logger *zap.Logger, middlewares middlewares.Middleware, config config.AppConfig) error {
	authController, err := controller.NewAuthController(goqu, logger, config)
	if err != nil {
		return err
	}
	v1.Post("/login", authController.DoAuth)

	if config.Kratos.IsEnabled {
		kratos := v1.Group("/kratos")
		kratos.Get("/auth", middlewares.Authenticated, authController.DoKratosAuth)
	}
	return nil
}

func setupUserController(v1 fiber.Router, goqu *goqu.Database, logger *zap.Logger, middlewares middlewares.Middleware, events *events.Events, pub *watermill.WatermillPublisher) error {
	userController, err := controller.NewUserController(goqu, logger, events, pub)
	if err != nil {
		return err
	}

	userRouter := v1.Group("/users")
	userRouter.Post("/", userController.CreateUser)
	userRouter.Get(fmt.Sprintf("/:%s", constants.ParamUid), middlewares.Authenticated, userController.GetUser)
	return nil
}

func healthCheckController(app *fiber.App, goqu *goqu.Database, logger *zap.Logger) error {
	healthController, err := controller.NewHealthController(goqu, logger)
	if err != nil {
		return err
	}

	healthz := app.Group("/healthz")
	healthz.Get("/", healthController.Overall)
	healthz.Get("/db", healthController.Db)
	return nil
}

func metricsController(app *fiber.App, db *goqu.Database, logger *zap.Logger, pMetrics *pMetrics.PrometheusMetrics) error {
	metricsController, err := controller.InitMetricsController(db, logger, pMetrics)
	if err != nil {
		return nil
	}

	app.Get("/metrics", metricsController.Metrics)
	return nil
}

func setupTitleController(v1 fiber.Router, goqu *goqu.Database, logger *zap.Logger, middlewares middlewares.Middleware, events *events.Events, pub *watermill.WatermillPublisher) error {
	titleController, err := controller.NewTitleController(goqu, logger, events, pub)
	if err != nil {
		return err
	}

	titleRouter := v1.Group("/titles")
	titleRouter.Post("/", titleController.Create)
	titleRouter.Get("/", titleController.List)
	titleRouter.Get(fmt.Sprintf("/:%s", constants.ParamTitleId), titleController.GetById)
	titleRouter.Delete(fmt.Sprintf("/:%s", constants.ParamTitleId), titleController.Delete)
	titleRouter.Put(fmt.Sprintf("/:%s", constants.ParamTitleId), titleController.Update)
	return nil
}

func setupCreditController(v1 fiber.Router, goqu *goqu.Database, logger *zap.Logger, middlewares *middlewares.TitleCreditCheckMiddleware, events *events.Events, pub *watermill.WatermillPublisher) error {
	creditController, err := controller.NewCreditController(goqu, logger, events, pub)
	if err != nil {
		return err
	}

	creditRouter := v1.Group(fmt.Sprintf("/titles/:%s/credits", constants.ParamTitleId))
	creditRouter.Post("/", middlewares.TitleExist, creditController.Create)
	creditRouter.Get("/", middlewares.TitleExist, creditController.ListByTitleId)
	creditRouter.Get(fmt.Sprintf("/:%s", constants.ParamCreditId), middlewares.CreditTitleVerification, creditController.GetById)
	creditRouter.Delete(fmt.Sprintf("/:%s", constants.ParamCreditId), middlewares.CreditTitleVerification, creditController.Delete)
	return nil
}
