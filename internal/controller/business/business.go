package business

import (
	"e-invoicing/internal/services/business"
	"e-invoicing/pkg/database"
	"e-invoicing/pkg/utility"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Db        *database.Database
	Validator *validator.Validate
	Logger    *utility.Logger
}

// @Summary      Get All Businesses
// @Description  Retrieve a list of all businesses in the system
// @Tags         Business
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Success      200 {object} models.Response "Businesses retrieved successfully"
// @Failure      400 {object} models.Response "Bad request"
// @Failure      401 {object} models.Response "Unauthorized"
// @Failure      500 {object} models.Response "Internal server error"
// @Router       /business [get]
func (base *Controller) GetAllBusiness(c *fiber.Ctx) error {
	businesses, err := business.GetAllBusinesses(base.Db.Postgresql.DB())
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", err.Error(), err, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "businesses gotten successfully", businesses)
	return c.Status(fiber.StatusOK).JSON(rd)
}

// @Summary      Get Business by ID
// @Description  Retrieve details of a specific business using its ID
// @Tags         Business
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        id   path      string  true  "Business ID" format(uuid)
// @Success      200 {object} models.Response "Business retrieved successfully"
// @Failure      400 {object} models.Response "Bad request"
// @Failure      401 {object} models.Response "Unauthorized"
// @Failure      404 {object} models.Response "Business not found"
// @Failure      500 {object} models.Response "Internal server error"
// @Router       /business/{id} [get]
func (base *Controller) GetBusinessByID(c *fiber.Ctx) error {
	id := c.Params("id")

	business, err := business.GetBusinessByID(base.Db.Postgresql.DB(), id)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusNotFound, "error", err.Error(), err, nil)
		return c.Status(http.StatusNotFound).JSON(rd)
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "business gotten successfully", business)
	return c.Status(http.StatusOK).JSON(rd)
}
