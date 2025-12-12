package invoice

import (
	"einvoice-access-point/external/firs_models"
	"einvoice-access-point/internal/services/invoice"
	"einvoice-access-point/pkg/utility"

	"github.com/gofiber/fiber/v2"
)

// // FirsWebhook godoc
// // @Summary      FIRS Webhook Receiver
// // @Description  Receives webhook events from FIRS (e.g., IRN status updates).
// // @Tags         Webhooks
// // @Accept       json
// // @Produce      json
// // @Param        payload  body      firs_models.FirsWebhookPayload  true  "FIRS Webhook Payload"
// // @Success      200      {object}  models.Response  "Webhook processed successfully"
// // @Failure      400      {object}  models.Response  "Bad request - invalid body or processing error"
// // @Failure      422      {object}  models.Response  "Validation failed"
// // @Router       /webhook/firs [post]
func (base *Controller) FirsWebhook(c *fiber.Ctx) error {
	var req firs_models.FirsWebhookPayload

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

	err = invoice.PrcoessFirsWebhook(req)
	if err != nil {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", err.Error(), err, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	base.Logger.Info("Webhook successfully reached for irn: %s", req.IRN)
	rd := utility.BuildSuccessResponse(fiber.StatusOK, "successful", nil)
	return c.Status(fiber.StatusOK).JSON(rd)
}
