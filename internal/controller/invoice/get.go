package invoice

import (
	"einvoice-access-point/external/firs_models"
	"einvoice-access-point/internal/services/invoice"
	"einvoice-access-point/pkg/utility"

	"github.com/gofiber/fiber/v2"
)

// ConfirmInvoice godoc
// @Summary Confirm Invoice
// @Description Confirms an invoice with IRN.
// @Tags Invoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param irn path string true "Invoice Reference Number (IRN)"
// @Success 200 {object} models.Response "Invoice confirmed successfully"
// @Failure 400 {object} models.Response "Bad request"
// @Router /invoice/confirm/{irn} [get]
func (base *Controller) ConfirmInvoice(c *fiber.Ctx) error {
	irn := c.Params("irn")
	if irn == "" {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", "irn is required", nil, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	respData, errDetails, err := invoice.ConfirmInvoice(irn)
	if err != nil {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", err.Error(), errDetails, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	base.Logger.Info("Invoice confirmed with irn successfully")
	rd := utility.BuildSuccessResponse(fiber.StatusOK, "Invoice confirmed with irn successfully", respData)
	return c.Status(fiber.StatusOK).JSON(rd)
}

// DownloadInvoice godoc
// @Summary Download Invoice
// @Description Downloads an invoice from FIRS using the IRN.
// @Tags Invoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param irn path string true "Invoice Reference Number (IRN)"
// @Success 200 {object} models.Response "Invoice downloaded successfully"
// @Failure 400 {object} models.Response "Bad request"
// @Router /invoice/download/{irn} [get]
func (base *Controller) DownloadInvoice(c *fiber.Ctx) error {
	irn := c.Params("irn")
	if irn == "" {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", "irn is required", nil, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	respData, errDetails, err := invoice.DownloadInvoice(irn)
	if err != nil {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", err.Error(), errDetails, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	base.Logger.Info("Invoice downloaded with irn successfully")
	rd := utility.BuildSuccessResponse(fiber.StatusOK, "Invoice downloaded with irn successfully", respData)
	return c.Status(fiber.StatusOK).JSON(rd)
}

// GetAllInvoicesByBusinessID godoc
// @Summary Get all invoices by business ID
// @Description Returns a list of invoices with minimal details for a business
// @Tags Internal Invoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param business_id path string true "Business ID" format(uuid)
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /invoice/business/{business_id} [get]
func (base *Controller) GetAllInvoicesByBusinessID(c *fiber.Ctx) error {
	businessID := c.Params("business_id")
	if businessID == "" {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", "business_id is required", nil, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	invoices, err := invoice.GetAllInvoicesByBusinessID(base.Db.Postgresql.DB(), businessID)
	if err != nil {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", err.Error(), err, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	rd := utility.BuildSuccessResponse(fiber.StatusOK, "Invoices fetched successfully", invoices)
	return c.Status(fiber.StatusOK).JSON(rd)
}

// GetInvoiceDetails godoc
// @Summary Get one invoice details
// @Description Returns full invoice details by business ID and invoice ID
// @Tags Internal Invoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param business_id path string true "Business ID" format(uuid)
// @Param invoice_id path string true "Invoice ID" format(uuid)
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /invoice/business/{business_id}/{invoice_id} [get]
func (base *Controller) GetInvoiceDetails(c *fiber.Ctx) error {
	businessID := c.Params("business_id")
	invoiceID := c.Params("invoice_id")

	if businessID == "" || invoiceID == "" {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", "business_id and invoice_id are required", nil, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	invoice, err := invoice.GetInvoiceDetails(base.Db.Postgresql.DB(), businessID, invoiceID)
	if err != nil {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", err.Error(), err, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	rd := utility.BuildSuccessResponse(fiber.StatusOK, "Invoice details fetched successfully", invoice)
	return c.Status(fiber.StatusOK).JSON(rd)
}

// CreateInvoice godoc
// @Summary Create a new Invoice
// @Description Upload a JSON invoice file and store it in DB
// @Tags Internal Invoice
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param file formData file true "Invoice JSON File"
// @Param business_id formData string true "Business ID"
// @Param invoice_number formData string true "Invoice Number"
// @Success 200 {object} models.Response "Invoice created successfully"
// @Failure 400 {object} models.Response "Bad request"
// @Router /invoice/create [post]
func (base *Controller) CreateInvoice(c *fiber.Ctx) error {

	businessID := c.FormValue("business_id")
	invoiceNumber := c.FormValue("invoice_number")

	if businessID == "" || invoiceNumber == "" {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", "business_id or invoice number is required", nil, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	file, err := c.FormFile("file")
	if err != nil {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", "invoice JSON file is required", nil, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	fileContent, err := file.Open()
	if err != nil {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", "failed to read file", nil, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}
	defer fileContent.Close()

	var payload firs_models.InvoiceRequest
	if err := utility.DecodeJSONWithDefaults(fileContent, &payload); err != nil {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", "invalid JSON format", nil, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	invoice, errDetails, err := invoice.CreateInvoice(base.Db.Postgresql.DB(), payload, invoiceNumber, businessID)
	if err != nil {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", err.Error(), errDetails, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	rd := utility.BuildSuccessResponse(fiber.StatusCreated, "Invoice created successfully", invoice)
	return c.Status(fiber.StatusCreated).JSON(rd)
}

// DeleteInvoice godoc
// @Summary Delete Invoice
// @Description Deletes an invoice by business_id and invoice_id
// @Tags Internal Invoice
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param business_id path string true "Business ID" format(uuid)
// @Param invoice_id path string true "Invoice ID" format(uuid)
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /invoice/business/{business_id}/{invoice_id} [delete]
func (base *Controller) DeleteInvoice(c *fiber.Ctx) error {
	businessID := c.Params("business_id")
	invoiceID := c.Params("invoice_id")

	if businessID == "" || invoiceID == "" {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", "business_id and invoice_id required", nil, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	if err := invoice.DeleteInvoice(base.Db.Postgresql.DB(), businessID, invoiceID); err != nil {
		rd := utility.BuildErrorResponse(fiber.StatusBadRequest, "error", err.Error(), err, nil)
		return c.Status(fiber.StatusBadRequest).JSON(rd)
	}

	rd := utility.BuildSuccessResponse(fiber.StatusOK, "Invoice deleted successfully", nil)
	return c.Status(fiber.StatusOK).JSON(rd)

}
