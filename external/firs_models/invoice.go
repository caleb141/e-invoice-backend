package firs_models

type InvoiceRequest struct {
	InvoiceNumber               string                 `json:"invoice_number"`
	BusinessID                  string                 `json:"business_id"`
	IRN                         *string                `json:"irn"`
	IssueDate                   string                 `json:"issue_date"`
	DueDate                     *string                `json:"due_date,omitempty"`
	IssueTime                   *string                `json:"issue_time,omitempty"`
	InvoiceTypeCode             string                 `json:"invoice_type_code"`
	PaymentStatus               *string                `json:"payment_status,omitempty"`
	Note                        *string                `json:"note,omitempty"`
	TaxPointDate                *string                `json:"tax_point_date,omitempty"`
	DocumentCurrencyCode        string                 `json:"document_currency_code"`
	TaxCurrencyCode             *string                `json:"tax_currency_code,omitempty"`
	AccountingCost              *string                `json:"accounting_cost,omitempty"`
	BuyerReference              *string                `json:"buyer_reference,omitempty"`
	InvoiceDeliveryPeriod       *InvoiceDeliveryPeriod `json:"invoice_delivery_period,omitempty"`
	OrderReference              *string                `json:"order_reference,omitempty"`
	BillingReference            []DocumentReference    `json:"billing_reference,omitempty"`
	DispatchDocumentReference   *DocumentReference     `json:"dispatch_document_reference,omitempty"`
	ReceiptDocumentReference    *DocumentReference     `json:"receipt_document_reference,omitempty"`
	OriginatorDocumentReference *DocumentReference     `json:"originator_document_reference,omitempty"`
	ContractDocumentReference   *DocumentReference     `json:"contract_document_reference,omitempty"`
	AdditionalDocumentReference []DocumentReference    `json:"_document_reference,omitempty"`
	AccountingSupplierParty     Party                  `json:"accounting_supplier_party"`
	AccountingCustomerParty     Party                  `json:"accounting_customer_party"`
	PayeeParty                  *Party                 `json:"payee_party,omitempty"`
	TaxRepresentativeParty      *Party                 `json:"tax_representative_party,omitempty"`
	ActualDeliveryDate          *string                `json:"actual_delivery_date,omitempty"`
	PaymentMeans                []PaymentMeans         `json:"payment_means,omitempty"`
	PaymentTermsNote            *string                `json:"payment_terms_note,omitempty"`
	AllowanceCharge             []AllowanceCharge      `json:"allowance_charge,omitempty"`
	TaxTotal                    []TaxTotal             `json:"tax_total,omitempty"`
	LegalMonetaryTotal          LegalMonetaryTotal     `json:"legal_monetary_total"`
	InvoiceLine                 []InvoiceLine          `json:"invoice_line"`
}

type InvoiceDeliveryPeriod struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type DocumentReference struct {
	IRN       string `json:"irn"`
	IssueDate string `json:"issue_date"`
}

type Party struct {
	PartyName           *string        `json:"party_name,omitempty"`
	TIN                 string         `json:"tin"`
	Email               string         `json:"email"`
	Telephone           *string        `json:"telephone,omitempty"`
	BusinessDescription *string        `json:"business_description,omitempty"`
	PostalAddress       *PostalAddress `json:"postal_address,omitempty"`
}

type PostalAddress struct {
	StreetName  string `json:"street_name,omitempty"`
	CityName    string `json:"city_name,omitempty"`
	PostalZone  string `json:"postal_zone,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

type PaymentMeans struct {
	PaymentMeansCode string `json:"payment_means_code"`
	PaymentDueDate   string `json:"payment_due_date"`
}

type AllowanceCharge struct {
	ChargeIndicator bool    `json:"charge_indicator"`
	Amount          float64 `json:"amount"`
}

type TaxTotal struct {
	TaxAmount   float64       `json:"tax_amount"`
	TaxSubtotal []TaxSubtotal `json:"tax_subtotal,omitempty"`
}

type TaxSubtotal struct {
	TaxableAmount float64     `json:"taxable_amount"`
	TaxAmount     float64     `json:"tax_amount"`
	TaxCategory   TaxCategory `json:"tax_category"`
}

type TaxCategory struct {
	ID      string  `json:"id"`
	Percent float64 `json:"percent"`
}

type LegalMonetaryTotal struct {
	LineExtensionAmount float64 `json:"line_extension_amount"`
	TaxExclusiveAmount  float64 `json:"tax_exclusive_amount"`
	TaxInclusiveAmount  float64 `json:"tax_inclusive_amount"`
	PayableAmount       float64 `json:"payable_amount"`
}

type InvoiceLine struct {
	HSNCode             string  `json:"hsn_code"`
	ProductCategory     string  `json:"product_category"`
	DiscountRate        float64 `json:"discount_rate"`
	DiscountAmount      float64 `json:"discount_amount"`
	FeeRate             float64 `json:"fee_rate"`
	FeeAmount           float64 `json:"fee_amount"`
	InvoicedQuantity    int     `json:"invoiced_quantity"`
	LineExtensionAmount float64 `json:"line_extension_amount"`
	Item                Item    `json:"item"`
	Price               Price   `json:"price"`
}

type Item struct {
	Name                      string  `json:"name"`
	Description               string  `json:"description"`
	SellersItemIdentification *string `json:"sellers_item_identification,omitempty"`
}

type Price struct {
	PriceAmount  float64 `json:"price_amount"`
	BaseQuantity int     `json:"base_quantity"`
	PriceUnit    string  `json:"price_unit"`
}

type IRNValidationRequest struct {
	InvoiceReference string `json:"invoice_reference" validate:"required"`
	BusinessID       string `json:"business_id" validate:"required"`
	IRN              string `json:"irn" validate:"required"`
}

type IRNValidationResponse struct {
	IRN       string `json:"IRN"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

type IRNSigningData struct {
	IRN         string `json:"irn"`
	Certificate string `json:"certificate"`
}

type IRNSigningResponse struct {
	EncryptedMessage string `json:"encrypted_message "`
	QrCodeImage      string `json:"qr_code_image"`
}

type IRNSigningRequestData struct {
	IRN string `json:"irn"`
}

type GenerateIRNRequestData struct {
	InvoiceNumber string `json:"invoice_number" validate:"required"`
}

type VerifyTinData struct {
	TIN string `json:"tin" validate:"required"`
}

type UpdateInvoice struct {
	PaymentStatus string  `json:"payment_status" validate:"required"`
	Reference     *string `json:"reference,omitempty"`
}

type FirsWebhookPayload struct {
	IRN     string `json:"irn" validate:"required"`
	Message string `json:"message" validate:"required"`
}
