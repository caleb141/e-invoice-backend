package v1

import (
	_ "e-invoicing/docs"
	"e-invoicing/pkg/config"
	"e-invoicing/pkg/database"
	"e-invoicing/pkg/middleware"
	"e-invoicing/pkg/utility"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func Setup(logger *utility.Logger, validator *validator.Validate, db *database.Database, keys *utility.CryptoKeys) *fiber.App {

	/////////////////////////////////////////////
	//Initiate Fiber App
	/////////////////////////////////////////////
	r := fiber.New(fiber.Config{
		Prefork:                 false,
		AppName:                 "eInvoice Firs Backend",
		JSONEncoder:             json.Marshal,
		JSONDecoder:             json.Unmarshal,
		ServerHeader:            "Golang Fiber",
		EnableTrustedProxyCheck: true,
		BodyLimit:               3 << 20,
		ErrorHandler:            middleware.GlobalErrorHandler,
	})

	r.Use(recover.New(recover.Config{
		EnableStackTrace: false,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			errMsg := "Unknown error"
			if err, ok := e.(error); ok {
				errMsg = err.Error()
			} else if msg, ok := e.(string); ok {
				errMsg = msg
			}

			rd := utility.BuildErrorResponse(fiber.StatusInternalServerError, "Internal Server Error", errMsg, nil, nil)

			_ = c.Status(fiber.StatusInternalServerError).JSON(rd)
		},
	}))

	r.Use(middleware.CORS())
	r.Use(middleware.Security())
	r.Use(middleware.Logger())
	r.Use(middleware.Metrics(config.GetConfig()))

	/////////////////////////////////////////////
	//General api route
	/////////////////////////////////////////////
	r.Get("/swagger/*", swagger.HandlerDefault)
	ApiVersion := "api/v1"
	/////////////////////////////////////////////
	//All Routes Registered
	/////////////////////////////////////////////

	HealthRoute(r, ApiVersion, validator, db, logger)
	AuthRoute(r, ApiVersion, validator, db, logger)
	EntityRoute(r, ApiVersion, validator, db, logger)
	BusinessRoute(r, ApiVersion, validator, db, logger)
	CallbackRoute(r, ApiVersion, validator, db, logger)
	InvoiceRoute(r, ApiVersion, validator, db, logger, keys)
	RegisterBaseRoutes(r, ApiVersion)

	return r
}
