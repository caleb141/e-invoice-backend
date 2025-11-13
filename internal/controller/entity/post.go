package entity

import (
	"e-invoicing/external/firs_models"
	"e-invoicing/internal/services/entity"
	"e-invoicing/pkg/utility"

	"github.com/gofiber/fiber/v2"
)

// @Summary      Verify TIN
// @Description  Verify a taxpayer identification number (TIN)
// @Tags         Entity
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        data  body      firs_models.VerifyTinData  true  "TIN verification request"
// @Success      200 {object} models.Response "TIN verified successfully"
// @Failure      400 {object} models.Response "Bad request"
// @Failure      401 {object} models.Response "Unauthorized"
// @Failure      422 {object} models.Response "Validation failed"
// @Failure      500 {object} models.Response "Internal server error"
// @Router       /entity/verify-tin [post]
func (base *Controller) VerifyTin(c *fiber.Ctx) error {
	var req firs_models.VerifyTinData

	err := c.BodyParser(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", "Failed to parse request body", err, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	err = base.Validator.Struct(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(fiber.StatusUnprocessableEntity, "error", "Validation failed", utility.ValidationResponse(err, base.Validator), nil)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(rd)
	}

	respData, errDetails, err := entity.VerifyTin(req.TIN)
	if err != nil {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", err.Error(), errDetails, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	base.Logger.Info("successfully")
	rd := utility.BuildSuccessResponse(fiber.StatusOK, "successfully", respData)
	return c.Status(fiber.StatusOK).JSON(rd)
}

// @Summary      Post VAT Payment
// @Description  Submit VAT payment details to FIRS
// @Tags         Entity
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        data  body      firs_models.FirsTransactionVatPayload  true  "VAT payment request payload"
// @Success      200 {object} models.Response "VAT payment processed successfully"
// @Failure      400 {object} models.Response "Bad request"
// @Failure      401 {object} models.Response "Unauthorized"
// @Failure      422 {object} models.Response "Validation failed"
// @Failure      500 {object} models.Response "Internal server error"
// @Router       /entity/vat-payment [post]
func (base *Controller) PostVatPayment(c *fiber.Ctx) error {
	var req firs_models.FirsTransactionVatPayload

	err := c.BodyParser(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", "Failed to parse request body", err, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	err = base.Validator.Struct(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(fiber.StatusUnprocessableEntity, "error", "Validation failed", utility.ValidationResponse(err, base.Validator), nil)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(rd)
	}

	respData, errDetails, err := entity.PostVatPayment(req)
	if err != nil {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", err.Error(), errDetails, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	base.Logger.Info("successfully")
	rd := utility.BuildSuccessResponse(fiber.StatusOK, "successfully", respData)
	return c.Status(fiber.StatusOK).JSON(rd)
}
