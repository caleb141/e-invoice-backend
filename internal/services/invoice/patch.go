package invoice

import (
	"e-invoicing/external/firs"
	"e-invoicing/external/firs_models"
	"fmt"
)

func UpdateInvoice(invoiceUpdate firs_models.UpdateInvoice, irn string) (*firs_models.FirsResponse, *string, error) {

	resp, err := firs.UpdateInvoice(invoiceUpdate, irn)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to validate irn: %w", err)
	}

	theResp, errDetails, err := firs.ParseFIRSAPIResponse(resp)
	if err != nil {
		return nil, errDetails, fmt.Errorf("failed to parse FIRS API response: %w", err)
	}

	//fmt.Println("IRN validation successful: ", theResp)
	return theResp, nil, nil
}
