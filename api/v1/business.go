package v1

import (
	"e-invoicing/internal/controller/business"
	"e-invoicing/pkg/database"
	"e-invoicing/pkg/middleware"
	"e-invoicing/pkg/utility"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BusinessRoute(app *fiber.App, ApiVersion string, validator *validator.Validate, db *database.Database, logger *utility.Logger) *fiber.App {
	businessController := business.Controller{Db: db, Validator: validator, Logger: logger}

	businessUrlSec := app.Group(fmt.Sprintf("%v/business", ApiVersion), middleware.Authorize(db.Postgresql.DB()))
	{
		businessUrlSec.Get("", businessController.GetAllBusiness)
		businessUrlSec.Get("/:id", businessController.GetBusinessByID)
	}

	return app
}
