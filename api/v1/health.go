package v1

import (
	"e-invoicing/internal/controller/health"
	"e-invoicing/pkg/database"
	"e-invoicing/pkg/utility"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HealthRoute(app *fiber.App, ApiVersion string, validator *validator.Validate, db *database.Database, logger *utility.Logger) *fiber.App {
	healthController := health.Controller{Db: db, Validator: validator, Logger: logger}

	healthGroup := app.Group(fmt.Sprintf("%v", ApiVersion))
	healthGroup.Post("/health", healthController.Post)
	healthGroup.Get("/health", healthController.Get)
	healthGroup.Get("/health/firs", healthController.FirsHealthCheck)

	return app
}
