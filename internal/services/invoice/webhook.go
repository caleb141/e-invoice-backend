package invoice

import (
	"e-invoicing/external/firs_models"
	"fmt"
)

func PrcoessFirsWebhook(payload firs_models.FirsWebhookPayload) error {

	_, theErr, err := TransmitConfirmInvoice(payload.IRN)
	if err == nil {
		return fmt.Errorf("failed to confirm transamitted invoice: %v - %v", *theErr, err)
	}

	fmt.Printf("irn: %s and message: %s", payload.IRN, payload.Message)
	return nil
}
