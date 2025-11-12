package entity

import (
	"einvoice-access-point/external/firs"
	"einvoice-access-point/external/firs_models"
	"fmt"
)

func VerifyTin(tin string) (*firs_models.FirsResponse, *string, error) {

	resp, err := firs.VerifyTin(tin)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to verify tin: %w", err)
	}

	theResp, errDetails, err := firs.ParseFIRSAPIResponse(resp)
	if err != nil {
		return nil, errDetails, fmt.Errorf("failed to parse FIRS API response: %w", err)
	}

	return theResp, nil, nil
}

func PostVatPayment(req firs_models.FirsTransactionVatPayload) (*firs_models.FirsResponse, *string, error) {

	resp, err := firs.PostVatPayment(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to post payment: %w", err)
	}

	theResp, errDetails, err := firs.ParseFIRSAPIResponse(resp)
	if err != nil {
		return nil, errDetails, fmt.Errorf("failed to parse FIRS API response: %w", err)
	}

	return theResp, nil, nil
}

func CreateParty(req firs_models.PartyRegistrationPayload) (*firs_models.FirsResponse, *string, error) {

	resp, err := firs.CreateParty(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create party: %w", err)
	}

	theResp, errDetails, err := firs.ParseFIRSAPIResponse(resp)
	if err != nil {
		return nil, errDetails, fmt.Errorf("failed to parse FIRS API response: %w", err)
	}

	return theResp, nil, nil
}
