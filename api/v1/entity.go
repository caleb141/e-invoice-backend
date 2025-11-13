package v1

import (
	"e-invoicing/internal/controller/entity"
	"e-invoicing/pkg/database"
	"e-invoicing/pkg/middleware"
	"e-invoicing/pkg/utility"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func EntityRoute(app *fiber.App, ApiVersion string, validator *validator.Validate, db *database.Database, logger *utility.Logger) *fiber.App {
	entityController := entity.Controller{Db: db, Validator: validator, Logger: logger}

	entityUrlSec := app.Group(fmt.Sprintf("%v/entity", ApiVersion), middleware.Authorize(db.Postgresql.DB()))
	{
		entityUrlSec.Get("", entityController.GetEntities)
		entityUrlSec.Get("/:entity_id", entityController.GetEntity)
		entityUrlSec.Post("/verify-tin", entityController.VerifyTin)
		entityUrlSec.Post("/vat-payment", entityController.PostVatPayment)
	}

	return app
}
