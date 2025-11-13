package invoice

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"e-invoicing/external/firs"
	"e-invoicing/external/firs_models"
	"e-invoicing/pkg/utility"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/png"
	"regexp"
	"strings"
	"time"

	qrcode "github.com/skip2/go-qrcode"
)

func PrepareIRN(irn string) string {
	timestamp := time.Now().UnixMilli()
	return fmt.Sprintf("%s.%d", irn, timestamp)
}

func GenerateIRNumber(invoiceNumber, serviceID string, timestamp time.Time) (string, error) {

	if !regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString(invoiceNumber) {
		return "", fmt.Errorf("invalid invoice number: only alphanumeric characters allowed")
	}

	if len(serviceID) != 8 || !regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString(serviceID) {
		return "", fmt.Errorf("invalid service ID: must be 8 alphanumeric characters")
	}

	dateString := timestamp.Format("20060102")

	irn := fmt.Sprintf("%s-%s-%s", invoiceNumber, serviceID, dateString)

	return irn, nil
}

func ValidateIRN(invoiceReq firs_models.IRNValidationRequest) (*firs_models.FirsResponse, *string, error) {

	resp, err := firs.ValidateIRN(invoiceReq)
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

func ValidateInvoice(invoiceReq firs_models.InvoiceRequest) (*firs_models.FirsResponse, *string, error) {

	resp, err := firs.ValidateInvoice(invoiceReq)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to validate invoice: %w", err)
	}

	theResp, errDetails, err := firs.ParseFIRSAPIResponse(resp)
	if err != nil {
		return nil, errDetails, fmt.Errorf("failed to parse FIRS API response: %w", err)
	}

	fmt.Println("Invoice validation successful: ", theResp)
	return theResp, nil, nil
}

func SignIRN(irn string, keys *utility.CryptoKeys) (*firs_models.IRNSigningResponse, error) {
	formattedIRN := PrepareIRN(irn)

	payload := firs_models.IRNSigningData{
		IRN:         formattedIRN,
		Certificate: keys.Certificate,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %v", err)
	}

	//encrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, keys.PublicKey, jsonData, nil)
	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, keys.PublicKey, jsonData)
	if err != nil {
		return nil, fmt.Errorf("encryption failed: %v", err)
	}

	base64Encrypted := base64.StdEncoding.EncodeToString(encrypted)

	qr, err := qrcode.New(base64Encrypted, qrcode.Medium)
	if err != nil {
		return nil, fmt.Errorf("failed to generate QR code: %v", err)
	}

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, qr.Image(256)); err != nil {
		return nil, fmt.Errorf("failed to encode QR code: %v", err)
	}

	base64QRImage := base64.StdEncoding.EncodeToString(buf.Bytes())

	theResp := &firs_models.IRNSigningResponse{
		EncryptedMessage: base64Encrypted,
		QrCodeImage:      base64QRImage,
	}

	//fmt.Printf("signed irn: %v", theResp)
	return theResp, nil
}

func SignInvoice(invoiceReq firs_models.InvoiceRequest) (*firs_models.FirsResponse, *string, error) {

	resp, err := firs.SignInvoice(invoiceReq)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to sign invoice: %w", err)
	}

	theResp, errDetails, err := firs.ParseFIRSAPIResponse(resp)
	if err != nil {
		return nil, errDetails, fmt.Errorf("failed to parse FIRS API response: %w", err)
	}

	fmt.Println("Invoice sign successful: ", theResp)
	return theResp, nil, nil
}

func GenerateIRN(invoiceNumber, serviceId string) (*string, error) {
	cleanInvoiceNumber := strings.ReplaceAll(invoiceNumber, "-", "")
	irn, err := GenerateIRNumber(cleanInvoiceNumber, serviceId, time.Now())
	if err != nil {
		return nil, err
	}
	return &irn, nil
}
