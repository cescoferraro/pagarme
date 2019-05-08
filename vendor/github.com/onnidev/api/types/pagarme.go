package types

import "time"

// ESTransactionSplitRules sdkfjn
type ESTransactionSplitRules struct {
	SplitRules string `json:"split_rules.recipient_id" bson:"split_rules.recipient_id"`
}

// ESTransactiosDate sdkfjn
type ESTransactiosDate struct {
	Date ESTransactionDateCreated `json:"date_created" bson:"date_created"`
}

// ESTransactionRange sdkfjn
type ESTransactionRange struct {
	Range ESTransactiosDate `json:"range" bson:"range"`
}

// ESTransactionTerm sdkfjn
type ESTransactionTerm struct {
	Term ESTransactionSplitRules `json:"term" bson:"term"`
}

// ESTransactionDateCreated sdkjfn
type ESTransactionDateCreated struct {
	Lte string `json:"lte" bson:"lte"`
	Gte string `json:"gte" bson:"gte"`
}

// ESTransactionFilter sdkfjn
type ESTransactionFilter struct {
	And []interface{} `json:"and" bson:"and"`
}

// ESTransactionFiltered sdkfjn
type ESTransactionFiltered struct {
	Filter ESTransactionFilter `json:"filter" bson:"filter"`
}

// ESFinanceQuery sdkfjn
type ESFinanceQuery struct {
	Filtered ESTransactionFiltered `json:"filtered" bson:"filtered"`
}

// PagarmeTransactionRequest type for the above middleware
type PagarmeTransactionRequest struct {
	Query ESFinanceQuery `json:"query" bson:"query"`
	Size  int            `json:"size" bson:"size"`
	From  int            `json:"from" bson:"from"`
}

// PagarmeSplitRule ksjdnf
type PagarmeSplitRule struct {
	Object      string `bson:"object" json:"object"`
	RecipientID string `json:"recipient_id" bson:"recipient_id"`
}

// PagarmeCardResponse type for the above middleware
type PagarmeCardResponse struct {
	Object         string `bson:"object" json:"object"`
	ID             string `bson:"id" json:"id"`
	DateCreated    string `bson:"date_created" json:"date_created"`
	DateUpdated    string `bson:"date_updated" json:"date_updated"`
	Brand          string `bson:"brand" json:"brand"`
	HolderName     string `bson:"holder_name" json:"holder_name"`
	FirstDigits    string `bson:"first_digits" json:"first_digits"`
	LastDigits     string `bson:"last_digits" json:"last_digits"`
	Country        string `bson:"country" json:"country"`
	FingerPrint    string `bson:"fingerprint" json:"fingerprint"`
	Customer       string `bson:"customer" json:"customer"`
	Valid          bool   `bson:"valid" json:"valid"`
	Expirationdate string `bson:"expiration_date" json:"expiration_date"`
}

// PagarmeMovementObject sdkfjn
type PagarmeMovementObject struct {
	Type   string `json:"type"`
	Status string `json:"status"`
	ID     int    `json:"id" bson:"id"`
}

// PagameOperation sdkjfn
type PagameOperation struct {
	Amount                int                   `json:"amount" bson:"amount"`
	Fee                   int                   `json:"fee" bson:"fee"`
	ID                    int                   `json:"id" bson:"id"`
	DateCreated           time.Time             `json:"date_created" bson:"date_created"`
	PagarmeMovementObject PagarmeMovementObject `json:"movement_object" bson:"movement_object"`
}

// SMSResponse sdkjnf
type SMSResponse struct {
	Password string `json:"password"`
}

// PagarMePayable sdkjfn
type PagarMePayable struct {
	ID                  int        `json:"id" bson:"id"`
	TransactionID       int        `json:"transaction_id" bson:"transaction_id"`
	Installment         int        `json:"installment" bson:"installment"`
	Object              string     `bson:"object" json:"object"`
	Status              string     `json:"status"`
	DateCreated         time.Time  `json:"date_created" bson:"date_created"`
	PaymentDate         time.Time  `json:"payment_date" bson:"payment_date"`
	AccrualDate         *time.Time `json:"accrual_date" bson:"accrual_date"`
	OriginalPaymentDate *time.Time `json:"original_payment_date" bson:"original_payment_date"`
	Type                string     `json:"type"`
	Amount              int        `json:"amount" bson:"amount"`
	PaymentMethod       string     `json:"payment_method"`
	Fee                 int        `json:"fee" bson:"fee"`
	AnticipationFee     int        `json:"antecipation_fee" bson:"antecipation_fee"`
	RecipientID         string     `json:"recipient_id" bson:"recipient_id"`
	SplitRuleID         string     `json:"split_rule_id" bson:"split_rule_id"`
	BulkAnticipationID  string     `json:"bulk_anticipation_id" bson:"bulk_anticipation_id"`
}

// PagarMeAntecipation sdkjfn
type PagarMeAntecipation struct {
	DateCreated     time.Time `json:"date_created" bson:"date_created"`
	DateUpdated     string    `bson:"date_updated" json:"date_updated"`
	Amount          int       `json:"amount" bson:"amount"`
	AnticipationFee int       `json:"antecipation_fee" bson:"antecipation_fee"`
	Fee             int       `json:"fee" bson:"fee"`
	ID              string    `json:"id" bson:"id"`
	Object          string    `bson:"object" json:"object"`
	PaymentDate     time.Time `json:"payment_date" bson:"payment_date"`
	Status          string    `json:"status"`
	TimeFrame       string    `json:"timeframe" bson:"timeframe"`
	Type            string    `json:"type" bson:"type"`
}

// PagarMeLimits sdkjfn
type PagarMeLimits struct {
	Maximum PagarMeAmount `json:"maximum" bson:"maximum"`
	Minimum PagarMeAmount `json:"minimum" bson:"minimum"`
}

// PagarMeAmount sdkfjn
type PagarMeAmount struct {
	Amount          int `json:"amount" bson:"amount"`
	AnticipationFee int `json:"antecipation_fee" bson:"antecipation_fee"`
	Fee             int `json:"fee" bson:"fee"`
}

// PagarMeTransfer sdkjfn
type PagarMeTransfer struct {
	Object      string             `bson:"object,omitempty" json:"object,omitempty"`
	ID          int                `json:"id,omitempty" bson:"id,omitempty"`
	Amount      int                `json:"amount,omitempty" bson:"amount,omitempty"`
	DateCreated *time.Time         `json:"date_created,omitempty" bson:"date_created,omitempty"`
	DateUpdated *time.Time         `bson:"date_updated,omitempty" json:"date_updated,omitempty"`
	Status      string             `json:"status,omitempty"`
	BankAccount PagarMeBankAccount `json:"bank_account,omitempty" bson:"bank_account,omitempty"`
}

// PagarMeReasons TODO: NEEDS COMMENT INFO
type PagarMeReasons struct {
	Type          string
	Message       string
	ParameterName *string
}

// PagarMeTransactionError TODO: NEEDS COMMENT INFO
type PagarMeTransactionError struct {
	Errors []PagarMeReasons `json:"errors" bson:"errors"`
	URL    string           `json:"url" bson:"url"`
	Method string           `json:"method" bson:"method"`
}

// PagarMeTransactionResponse sdfn
type PagarMeTransactionResponse struct {
	Object string `bson:"object" json:"object"`
	ID     int    `json:"id" bson:"id"`

	// Valores possíveis: processing, authorized, paid, refunded, waiting_payment, pending_refund, refused .
	Status string `json:"status" bson:"status"`
	// Valores possíveis: acquirer, antifraud, internal_error, no_acquirer, acquirer_timeout
	RefuseReason string `json:"refuse_reason" bson:"refuse_reason"`
	// Valores possíveis: acquirer, antifraud, internal_error, no_acquirer, acquirer_timeout
	StatusReason string `json:"status_reason" bson:"status_reason"`

	// Valores possíveis: development (em ambiente de testes), pagarme (adquirente Pagar.me), stone, cielo, rede.
	AcquirerName         string `json:"acquirer_name" bson:"acquirer_name"`
	AcquirerID           string `json:"acquirer_id" bson:"acquirer_id"`
	AcquirerResponseCode string `json:"acquirer_response_code" bson:"acquirer_response_code"`
	AuthorizationCode    string `json:"authorization_code" bson:"authorization_code"`

	SoftDescriptor string `json:"soft_descriptor" bson:"soft_descriptor"`
	TID            int    `json:"tid" bson:"tid"`
	NSU            int    `json:"nsu" bson:"nsu"`

	DateCreated *time.Time `json:"date_created,omitempty" bson:"date_created,omitempty"`
	DateUpdated *time.Time `bson:"date_updated,omitempty" json:"date_updated,omitempty"`
	Amount      int        `json:"amount" bson:"amount"`

	AuthorizedAmount     int                 `json:"authorized_amount" bson:"authorized_amount"`
	PaidAmount           int                 `json:"paid_amount" bson:"paid_amount"`
	RefundedAmount       int                 `json:"refunded_amount" bson:"refunded_amount"`
	Installments         int                 `json:"installments" bson:"installments"`
	Cost                 int                 `json:"cost" bson:"cost"`
	CardHolderName       string              `json:"card_holder_name" bson:"card_holder_name"`
	CardLastDigits       string              `json:"card_last_digits" bson:"card_last_digits"`
	CardFirstDigits      string              `json:"card_first_digits" bson:"card_first_digits"`
	CardBrand            string              `json:"card_brand" bson:"card_brand"`
	CardPinMode          string              `json:"card_pin_mode" bson:"card_pin_mode"`
	PostBackURL          string              `json:"postback_url" bson:"postback_url"`
	PaymentMethod        string              `json:"payment_method" bson:"payment_method"`
	CaptureMethod        string              `json:"capture_method" bson:"capture_method"`
	AntiFraudScore       string              `json:"antifraud_score" bson:"antifraud_score"`
	AntiFraudMetadata    AntiFraudMetadata   `json:"antifraud_metadata" bson:"antifraud_metadata"`
	BoletoURL            string              `json:"boleto_url" bson:"boleto_url"`
	BoletoBarCode        string              `json:"boleto_barcode" bson:"boleto_barcode"`
	BoletoExpirationDate string              `json:"boleto_expiration_date" bson:"boleto_expiration_date"`
	Referer              string              `json:"referer" bson:"referer"`
	IP                   string              `json:"ip" bson:"ip"`
	SubscriptionID       int                 `json:"subscription_id" bson:"subscription_id"`
	Customer             PagarMeCustomer     `json:"customer" bson:"customer"`
	MetaData             MetaData            `json:"metadata" bson:"metadata"`
	SplitRules           []SplitRuleResponse `json:"split_rules" bson:"split_rules"`
	Session              MetaData            `json:"session" bson:"session"`
	Items                []PagarMeItens      `json:"items" bson:"items"`
	Address              PagarMeBilling      `json:"address" bson:"address"`
	Billing              PagarMeBilling      `json:"billing" bson:"billing"`
	Shipping             PagarMeBilling      `json:"shipping" bson:"shipping"`
	Documents            []DocumentFull      `json:"documents" bson:"documents"`
}

// AntiFraudMetadata sdkfjn
type AntiFraudMetadata struct {
}

// PagarMeItens sdkfjn
type PagarMeItens struct {
	ID        int    `json:"id" bson:"id"`
	Title     string `json:"title" bson:"title"`
	UnitPrice int    `json:"unit_price" bson:"unit_price"`
	Quantity  int    `json:"quantity" bson:"quantity"`
	Tangible  bool   `json:"tangible" bson:"tangible"`
	Category  string `json:"category" bson:"category"`
	Venue     string `json:"venue" bson:"venue"`
	Date      string `json:"date" bson:"date"`
}

// PagarMeBilling struct
type PagarMeBilling struct {
	Name    string         `json:"name" bson:"name"`
	Address PagarMeAddress `json:"address" bson:"address"`
}

// PagarMeAddress struct
type PagarMeAddress struct {
	Street        string `json:"street" bson:"street"`
	Country       string `json:"country" bson:"country"`
	City          string `json:"city" bson:"city"`
	State         string `json:"state" bson:"state"`
	Number        string `json:"number" bson:"number"`
	ZipCode       string `json:"zipcode" bson:"zipcode"`
	Neighborhood  string `json:"neighborhood" bson:"neighborhood"`
	Complementary string `json:"complementary" bson:"complementary"`
}

// PagarMeRecipient sdkjfn
type PagarMeRecipient struct {
	Object                         string             `bson:"object,omitempty" json:"object,omitempty"`
	ID                             string             `json:"id,omitempty" bson:"id,omitempty"`
	TransferEnabled                bool               `json:"transfer_enabled,omitempty" bson:"transfer_enabled,omitempty"`
	AutomaticAntecipationEnabled   bool               `json:"automatic_anticipation_enabled,omitempty" bson:"automatic_anticipation_enabled,omitempty"`
	AutomaticAntecipationType      string             `json:"automatic_anticipation_type,omitempty" bson:"automatic_anticipation_type,omitempty"`
	AutomaticAntecipation1025Delay int                `json:"automatic_anticipation_1025_delay,omitempty" bson:"automatic_anticipation_1025_delay,omitempty"`
	AntecipableVolumePercentage    int                `json:"anticipatable_volume_percentage,omitempty" bson:"anticipatable_volume_percentage,omitempty"`
	AutomaticAntecipationDays      string             `json:"automatic_anticipation_days,omitempty" bson:"automatic_anticipation_days,omitempty"`
	LastTransfer                   *time.Time         `json:"last_transfer,omitempty" bson:"last_transfer,omitempty"`
	TransferInterval               string             `json:"transfer_interval,omitempty" bson:"transfer_interval,omitempty"`
	PostBackURL                    string             `json:"postback_url,omitempty" bson:"postback_url,omitempty"`
	TransferDay                    int                `json:"transfer_day,omitempty" bson:"transfer_day,omitempty"`
	DateCreated                    *time.Time         `json:"date_created,omitempty" bson:"date_created,omitempty"`
	DateUpdated                    *time.Time         `bson:"date_updated,omitempty" json:"date_updated,omitempty"`
	Status                         string             `json:"status,omitempty"`
	StatusReason                   string             `json:"status_reason,omitempty" bson:"status_reason,omitempty"`
	BankAccount                    PagarMeBankAccount `json:"bank_account,omitempty" bson:"bank_account,omitempty"`
}

// PagarMeBankAccount sdkjfn
type PagarMeBankAccount struct {
	Object             string    `bson:"object,omitempty" json:"object,omitempty"`
	ID                 int       `json:"id,omitempty" bson:"id,omitempty"`
	BankCode           string    `json:"bank_code" bson:"bank_code"`
	Conta              string    `json:"conta" bson:"conta"`
	ContaDV            string    `json:"conta_dv" bson:"conta_dv"`
	Agencia            string    `json:"agencia" bson:"agencia"`
	AgenciaDV          *string   `json:"agencia_dv" bson:"agencia_dv"`
	DocumentType       string    `json:"document_type,omitempty" bson:"document_type,omitempty"`
	DocumentNumber     string    `json:"document_number" bson:"document_number"`
	Type               string    `json:"type,omitempty" bson:"type,omitempty"`
	LegalName          string    `json:"legal_name" bson:"legal_name"`
	ChargeTransferFees bool      `json:"charge_transfer_fees,omitempty" bson:"charge_transfer_fees,omitempty"`
	DateCreated        time.Time `json:"date_created,omitempty" bson:"date_created,omitempty"`
}

// PagarMeAntecipationPostRequest type for the above middleware
type PagarMeAntecipationPostRequest struct {
	APIKey          string `bson:"api_key" json:"api_key"`
	Timeframe       string `json:"timeframe" bson:"timeframe"`
	Build           bool   `json:"build" bson:"build"`
	PaymentDay      int    `json:"payment_date" bson:"payment_date"`
	RequestedAmount int32  `json:"requested_amount" bson:"requested_amount"`
}

// PagarMeAntecipationResponse type for the above middleware
type PagarMeAntecipationResponse struct {
	Amount          int       `json:"amount" bson:"amount"`
	AnticipationFee int       `json:"anticipation_fee" bson:"anticipation_fee"`
	DateCreated     time.Time `bson:"date_created" json:"date_created"`
	DateUpdated     time.Time `bson:"date_updated" json:"date_updated"`
	PaymentDate     time.Time `json:"payment_date" bson:"payment_date"`
	Fee             int       `json:"fee" bson:"fee"`
	ID              string    `json:"id" bson:"id"`
	Object          string    `bson:"object" json:"object"`
	Type            string    `json:"type"`
	Status          string    `json:"status" bson:"status"`
	Timeframe       string    `json:"timeframe" bson:"timeframe"`
}

// PagarMeRecipientPostRequest type for the above middleware
type PagarMeRecipientPostRequest struct {
	AutomaticAntecipationsEnabled string                          `json:"automatic_anticipation_enabled" bson:"automatic_anticipation_enabled"`
	AntecipatableVolumePercentage string                          `json:"anticipatable_volume_percentage" bson:"anticipatable_volume_percentage"`
	TransferDay                   string                          `json:"transfer_day" bson:"transfer_enabled"`
	TransferEnabled               string                          `json:"transfer_enabled" bson:"transfer_enabled"`
	TransferInterval              string                          `json:"transfer_interval" bson:"transfer_interval"`
	PostBackURL                   string                          `json:"postback_url" bson:"postback_url"`
	BankAccount                   PagarMeRecipientPostBankAccount `json:"bank_account" bson:"bank_account"`
}

// PagarMeRecipientPostBankAccount sdkjfn
type PagarMeRecipientPostBankAccount struct {
	DocumentType   string  `json:"document_type" bson:"document_type"`
	LegalName      string  `json:"legal_name" bson:"legal_name"`
	BankCode       string  `json:"bank_code" bson:"bank_code"`
	Branch         string  `json:"agencia" bson:"agencia"`
	BranchVC       *string `json:"agencia_dv,omitempty" bson:"agencia_dv,omitempty"`
	Account        string  `json:"conta" bson:"conta"`
	Type           string  `json:"type" bson:"type"`
	AccountVC      string  `json:"conta_dv" bson:"conta_dv"`
	DocumentNumber string  `json:"document_number" bson:"document_number"`
}

// PagarMeTransactionRefundRequest TODO: NEEDS COMMENT INFO
type PagarMeTransactionRefundRequest struct {
	Amount     float64     `json:"amount" bson:"amount"`
	APIKey     string      `bson:"api_key" json:"api_key"`
	Async      string      `json:"async" bson:"async"`
	SplitRules []SplitRule `json:"split_rules" bson:"split_rules"`
}

// PagarMeTransactionRequest TODO: NEEDS COMMENT INFO
type PagarMeTransactionRequest struct {
	Amount          float64         `json:"amount" bson:"amount"`
	PaymentMethod   string          `json:"payment_method"`
	APIKey          string          `bson:"api_key" json:"api_key"`
	Async           string          `json:"async" bson:"async"`
	CardID          string          `json:"card_id" bson:"card_id"`
	SoftDescriptor  string          `json:"soft_descriptor" bson:"soft_descriptor"`
	Capture         string          `json:"capture" bson:"capture"`
	PagarMeCustomer PagarMeCustomer `json:"customer" bson:"customer"`
	SplitRules      []SplitRule     `json:"split_rules" bson:"split_rules"`
	MetaData        MetaData        `json:"metadata" bson:"metadata"`
}

// MetaData sdfjknsdfk
type MetaData struct {
	ClubID     string `json:"clubId" bson:"clubId,omitempty"`
	ClubName   string `json:"clubName" bson:"clubName"`
	CustomerID string `json:"customerId" bson:"customerId"`
	PartyID    string `json:"partyId" bson:"partyId"`
	PartyName  string `json:"partyName" bson:"partyName"`
}

// PagarMeCustomer TODO: NEEDS COMMENT INFO
type PagarMeCustomer struct {
	Type           string     `json:"type" bson:"type"`
	Country        string     `json:"country" bson:"country"`
	Name           string     `json:"name" bson:"name"`
	Email          string     `json:"email,omitempty" bson:"email,omitempty"`
	DocumentNumber string     `json:"documentNumber" bson:"documentNumber"`
	Documents      []Document `json:"documents" bson:"documents"`
}

// SplitRuleResponse TODO: NEEDS COMMENT INFO
type SplitRuleResponse struct {
	ID                  string `json:"id" bson:"id"`
	Amount              int    `json:"amount" bson:"amount"`
	RecipientID         string `json:"recipient_id" bson:"recipient_id"`
	Liable              bool   `json:"liable" bson:"liable"`
	ChargeProcessingFee bool   `json:"charge_processing_fee" bson:"charge_processing_fee"`
	ChargeRemainderFee  bool   `json:"charge_remainder_fee" bson:"charge_remainder_fee"`
}

// SplitRule TODO: NEEDS COMMENT INFO
type SplitRule struct {
	ID                  string `json:"id,omitempty" bson:"id,omitempty"`
	Amount              string `json:"amount" bson:"amount"`
	RecipientID         string `json:"recipient_id" bson:"recipient_id"`
	Liable              string `json:"liable" bson:"liable"`
	ChargeProcessingFee string `json:"charge_processing_fee" bson:"charge_processing_fee"`
	ChargeRemainderFee  string `json:"charge_remainder_fee" bson:"charge_remainder_fee"`
}

// Document TODO: NEEDS COMMENT INFO
type Document struct {
	Type string `json:"type" bson:"type"`
}

// DocumentFull TODO: NEEDS COMMENT INFO
type DocumentFull struct {
	Type string `json:"type" bson:"type"`
	Name string `json:"name" bson:"name"`
}
