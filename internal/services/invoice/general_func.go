package invoice

import (
	"e-invoicing/external/firs_models"
	repository "e-invoicing/internal/repository/invoice"
	inst "e-invoicing/pkg/dbinit"
	"e-invoicing/pkg/models"
	"fmt"

	"gorm.io/gorm"
)

func FirsAllInOneProcess(payload firs_models.InvoiceRequest, invoiceModel *models.Invoice, db *gorm.DB) error {

	pdb := inst.InitDB(db, true)

	_, theErr, err := ValidateInvoice(payload)
	if err != nil {
		_ = repository.UpdateInvoiceStatus(pdb, invoiceModel, models.StatusValidatedInvoice, "failed")
		return fmt.Errorf("failed to validate invoice: %v - %v", *theErr, err)
	}

	err = repository.UpdateInvoiceStatus(pdb, invoiceModel, models.StatusValidatedInvoice, "success")
	if err != nil {
		return fmt.Errorf("failed to update invoice status: %v", err)
	}

	_, theErr, err = SignInvoice(payload)
	if err != nil {
		_ = repository.UpdateInvoiceStatus(pdb, invoiceModel, models.StatusSignedInvoice, "failed")
		return fmt.Errorf("failed to sign invoice: %v - %v", *theErr, err)
	}
	err = repository.UpdateInvoiceStatus(pdb, invoiceModel, models.StatusSignedInvoice, "success")
	if err != nil {
		return fmt.Errorf("failed to update invoice status: %v", err)
	}

	_, theErr, err = TransmitInvoice(payload.IRN)
	if err != nil {
		_ = repository.UpdateInvoiceStatus(pdb, invoiceModel, models.StatusTransmitted, "failed")
		return fmt.Errorf("failed to transmit invoice: %v - %v", *theErr, err)
	}
	err = repository.UpdateInvoiceStatus(pdb, invoiceModel, models.StatusTransmitted, "success")
	if err != nil {
		return fmt.Errorf("failed to update invoice status: %v", err)
	}

	_, theErr, err = TransmitConfirmInvoice(payload.IRN)
	if err != nil {
		_ = repository.UpdateInvoiceStatus(pdb, invoiceModel, models.StatusConfirmed, "failed")
		return fmt.Errorf("failed to confirm transmit invoice: %v - %v", *theErr, err)
	}
	err = repository.UpdateInvoiceStatus(pdb, invoiceModel, models.StatusConfirmed, "success")
	if err != nil {
		return fmt.Errorf("failed to update invoice status: %v", err)
	}

	confirmInvoiceResp, theErr, err := ConfirmInvoice(payload.IRN)
	if err != nil {
		return fmt.Errorf("failed to confirm invoice: %v - %v", *theErr, err)
	}

	if confirmInvoiceResp.Code != 200 {
		return fmt.Errorf("failed to confirm invoice, didnt get 200 or delivered is false")
	}

	return nil
}
