package visa

import "encoding/json"

var PULL_FUNDS_TRANSACTIONS_URL = API_URL + "/visadirect/fundstransfer/v1/pullfundstransactions/"
var PULL_MULTI_FUNDS_TRANSACTIONS_URL = API_URL + "/visadirect/fundstransfer/v1/multipullfundstransactions/"
var PUSH_FUNDS_TRANSACTIONS_URL = API_URL + "/visadirect/fundstransfer/v1/pushfundstransactions/"
var PUSH_MULTI_FUNDS_TRANSACTIONS_URL = API_URL + "/visadirect/fundstransfer/v1/multipushfundstransactions/"
var REVERSE_FUNDS_TRANSACTIONS_URL = API_URL + "/visadirect/fundstransfer/v1/reversefundstransactions/"

type PullFundsTransactionRequest struct {
	SystemsTraceAuditNumber       int                       `json:"systemsTraceAuditNumber"`                 // required, 6
	RetrievalReferenceNumber      string                    `json:"retrievalReferenceNumber"`                // ydddhhnnnnnn(numeric characters only), Length: 12
	LocalTransactionDateTime      string                    `json:"localTransactionDateTime"`                // RFC3339. dateTime | YYYY-MM-DDThh:mm:ss. The date and time you submit the transaction
	AcquiringBin                  int                       `json:"acquiringBin"`                            // integer | positive, Length: 6 - 11
	AcquirerCountryCode           int                       `json:"acquirerCountryCode"`                     // integer | Length: 3
	SenderPrimaryAccountNumber    string                    `json:"senderPrimaryAccountNumber"`              // string | Length: 13 - 19
	SenderCardExpiryDate          string                    `json:"senderCardExpiryDate"`                    // string | YYYY-MM
	SenderCurrencyCode            string                    `json:"senderCurrencyCode"`                      // string | Length: 3
	Amount                        float64                   `json:"amount,omitempty"`                        // Optional: decimal | Length: totalDigits 12, fractionDigits 3 (minimum value is 0)
	Surcharge                     float64                   `json:"surcharge,omitempty"`                     // Optional: decimal | Length: totalDigits 12, fractionDigits 3(minimum value is 0)
	Cavv                          string                    `json:"cavv"`                                    // string | Length:40
	ForeignExchangeFeeTransaction float64                   `json:"foreignExchangeFeeTransaction,omitempty"` // Optional: decimal | Length: totalDigits 12, fractionDigits 3 (minimum value is 0)
	BusinessApplicationId         string                    `json:"businessApplicationId"`                   // string | Length: 2
	MerchantCategoryCode          int                       `json:"merchantCategoryCode,omitempty"`          // Conditional: integer | Length: total 4 digits
	CardAcceptor                  CardAcceptor              `json:"cardAcceptor"`                            // Object
	MagneticStripeData            *MagneticStripeData       `json:"magneticStripeData,omitempty"`            // Optional: Object
	PointOfServiceData            *PointOfServiceData       `json:"pointOfServiceData,omitempty"`            // Conditional: Object
	PointOfServiceCapability      *PointOfServiceCapability `json:"pointOfServiceCapability,omitempty"`      // Conditional: Object
	PinData                       *PinData                  `json:"pinData,omitempty"`                       // Conditional: Object
	FeeProgramIndicator           string                    `json:"feeProgramIndicator,omitempty"`           // Optional: string | Length:3
}

type PushFundsTransactionRequest struct {
	SystemsTraceAuditNumber       int                       `json:"systemsTraceAuditNumber"`            // required, 6
	RetrievalReferenceNumber      string                    `json:"retrievalReferenceNumber"`           // ydddhhnnnnnn(numeric characters only), Length: 12
	LocalTransactionDateTime      string                    `json:"localTransactionDateTime"`           // RFC3339. dateTime | YYYY-MM-DDThh:mm:ss. The date and time you submit the transaction
	AcquiringBin                  int                       `json:"acquiringBin"`                       // integer | positive, Length: 6 - 11
	AcquirerCountryCode           int                       `json:"acquirerCountryCode"`                // integer | Length: 3
	SenderAccountNumber           string                    `json:"senderAccountNumber,omitempty"`      // Conditional: string | Length: 0 - 34
	SenderAddress                 string                    `json:"senderAddress,omitempty"`            // Conditional: string | Length: 1 to 35
	SenderCity                    string                    `json:"senderCity,omitempty"`               // Conditional: string | Length: 1 to 25
	SenderStateCode               string                    `json:"senderStateCode,omitempty"`          // Optional: string | Length: 2
	SenderCountryCode             string                    `json:"senderCountryCode,omitempty"`        // Optional: string | Length: 2 or 33
	SenderName                    string                    `json:"senderName,omitempty"`               // Optional: string | Length: 1 to 30
	SenderReference               string                    `json:"senderReference,omitempty"`          // Optional: string | only alphabets (a-z, A-Z) and/or numbers (0-9) allowed , max: 16 characters
	SenderDateOfBirth             string                    `json:"senderDateOfBirth,omitempty"`        // Optional: string | YYYY-MM-DD
	RecipientName                 string                    `json:"recipientName,omitempty"`            // Conditional: string | Length: minimum 1, maximum 30
	RecipientPrimaryAccountNumber string                    `json:"recipientPrimaryAccountNumber"`      // string | Length: 13 - 19
	TransactionIdentifier         int                       `json:"transactionIdentifier"`              // integer | positive, Length: 15
	TransactionCurrencyCode       string                    `json:"transactionCurrencyCode"`            // string | Length: 3
	SourceOfFundsCode             string                    `json:"sourceOfFundsCode,omitempty"`        // Conditional: string | Length: 2
	Amount                        float64                   `json:"amount,omitempty"`                   // Optional: decimal | Length: totalDigits 12, fractionDigits 3 (minimum value is 0)
	BusinessApplicationId         string                    `json:"businessApplicationId"`              // string | Length: 2
	MerchantCategoryCode          int                       `json:"merchantCategoryCode,omitempty"`     // Conditional: integer | Length: total 4 digits
	CardAcceptor                  CardAcceptor              `json:"cardAcceptor"`                       // Object
	MagneticStripeData            *MagneticStripeData       `json:"magneticStripeData,omitempty"`       // Optional: Object
	PointOfServiceData            *PointOfServiceData       `json:"pointOfServiceData,omitempty"`       // Conditional: Object
	PointOfServiceCapability      *PointOfServiceCapability `json:"pointOfServiceCapability,omitempty"` // Conditional: Object
	PinData                       *PinData                  `json:"pinData,omitempty"`                  // Conditional: Object
	FeeProgramIndicator           string                    `json:"feeProgramIndicator,omitempty"`      // Optional: string | Length:3
}

type ReverseFundsTransactionRequest struct {
	SystemsTraceAuditNumber       int                         `json:"systemsTraceAuditNumber"`                 // required, 6
	RetrievalReferenceNumber      string                      `json:"retrievalReferenceNumber"`                // ydddhhnnnnnn(numeric characters only), Length: 12
	LocalTransactionDateTime      string                      `json:"localTransactionDateTime"`                // RFC3339. dateTime | YYYY-MM-DDThh:mm:ss. The date and time you submit the transaction
	AcquiringBin                  int                         `json:"acquiringBin"`                            // integer | positive, Length: 6 - 11
	AcquirerCountryCode           int                         `json:"acquirerCountryCode"`                     // integer | Length: 3
	SenderPrimaryAccountNumber    string                      `json:"senderPrimaryAccountNumber,omitempty"`    // Conditional: string | Length: 0 - 34
	SenderCardExpiryDate          string                      `json:"senderCardExpiryDate,omitempty"`          // Conditional: string | Length: 1 to 35
	SenderCurrencyCode            string                      `json:"senderCurrencyCode,omitempty"`            // Conditional: string | Length: 1 to 25
	TransactionIdentifier         int                         `json:"transactionIdentifier"`                   // integer | positive, Length: 15
	Amount                        float64                     `json:"amount,omitempty"`                        // Optional: decimal | Length: totalDigits 12, fractionDigits 3 (minimum value is 0)
	Surcharge                     float64                     `json:"surcharge,omitempty"`                     // Optional: decimal | Length: totalDigits 12, fractionDigits 3(minimum value is 0)
	ForeignExchangeFeeTransaction float64                     `json:"foreignExchangeFeeTransaction,omitempty"` // Optional: decimal | Length: totalDigits 12, fractionDigits 3 (minimum value is 0)
	OriginalDataElements          ReverseOriginalDataElements `json:"originalDataElements"`                    // Object
	CardAcceptor                  CardAcceptor                `json:"cardAcceptor"`                            // Object
	PointOfServiceData            *PointOfServiceData         `json:"pointOfServiceData,omitempty"`            // Conditional: Object
	PointOfServiceCapability      *PointOfServiceCapability   `json:"pointOfServiceCapability,omitempty"`      // Conditional: Object
	FeeProgramIndicator           string                      `json:"feeProgramIndicator,omitempty"`           // Optional: string | Length:3
}

type ReverseOriginalDataElements struct {
	ApprovalCode            string `json:"approvalCode"`            // string | Length: 6
	SystemsTraceAuditNumber int    `json:"systemsTraceAuditNumber"` // integer | required, 6
	TransmissionDateTime    string `json:"transmissionDateTime"`    // dateTime | YYYY-MM-DDThh:mm:ss
	AcquiringBin            int    `json:"acquiringBin"`            // integer | positive, Length: 6 - 11
}

type PullFundsTransactionRequestMulti struct {
	LocalTransactionDateTime string                                 `json:"localTransactionDateTime"`       // RFC3339. dateTime | YYYY-MM-DDThh:mm:ss. The date and time you submit the transaction
	AcquiringBin             int                                    `json:"acquiringBin"`                   // integer | positive, Length: 6 - 11
	AcquirerCountryCode      int                                    `json:"acquirerCountryCode"`            // integer | Length: 3
	BusinessApplicationId    string                                 `json:"businessApplicationId"`          // string | Length: 2
	MerchantCategoryCode     int                                    `json:"merchantCategoryCode,omitempty"` // Conditional: integer | Length: total 4 digits
	Request                  []PullFundsTransactionRequestMultiData `json:"request"`
	FeeProgramIndicator      string                                 `json:"feeProgramIndicator,omitempty"` // Optional: string | Length:3
}

type PushFundsTransactionRequestMulti struct {
	LocalTransactionDateTime string                                 `json:"localTransactionDateTime"`       // RFC3339. dateTime | YYYY-MM-DDThh:mm:ss. The date and time you submit the transaction
	AcquiringBin             int                                    `json:"acquiringBin"`                   // integer | positive, Length: 6 - 11
	AcquirerCountryCode      int                                    `json:"acquirerCountryCode"`            // integer | Length: 3
	BusinessApplicationId    string                                 `json:"businessApplicationId"`          // string | Length: 2
	MerchantCategoryCode     int                                    `json:"merchantCategoryCode,omitempty"` // Conditional: integer | Length: total 4 digits
	Request                  []PushFundsTransactionRequestMultiData `json:"request"`
	FeeProgramIndicator      string                                 `json:"feeProgramIndicator,omitempty"` // Optional: string | Length:3
}

type PullFundsTransactionRequestMultiResponse struct {
	LocalTransactionDateTime string                                 `json:"localTransactionDateTime"`       // RFC3339. dateTime | YYYY-MM-DDThh:mm:ss. The date and time you submit the transaction
	AcquiringBin             int                                    `json:"acquiringBin"`                   // integer | positive, Length: 6 - 11
	AcquirerCountryCode      int                                    `json:"acquirerCountryCode"`            // integer | Length: 3
	BusinessApplicationId    string                                 `json:"businessApplicationId"`          // string | Length: 2
	MerchantCategoryCode     int                                    `json:"merchantCategoryCode,omitempty"` // Conditional: integer | Length: total 4 digits
	Response                 []PullFundsTransactionRequestMultiData `json:"response"`
	FeeProgramIndicator      string                                 `json:"feeProgramIndicator,omitempty"` // Optional: string | Length:3
}

type PushFundsTransactionRequestMultiResponse struct {
	LocalTransactionDateTime string                                 `json:"localTransactionDateTime"`       // RFC3339. dateTime | YYYY-MM-DDThh:mm:ss. The date and time you submit the transaction
	AcquiringBin             int                                    `json:"acquiringBin"`                   // integer | positive, Length: 6 - 11
	AcquirerCountryCode      int                                    `json:"acquirerCountryCode"`            // integer | Length: 3
	BusinessApplicationId    string                                 `json:"businessApplicationId"`          // string | Length: 2
	MerchantCategoryCode     int                                    `json:"merchantCategoryCode,omitempty"` // Conditional: integer | Length: total 4 digits
	Response                 []PushFundsTransactionRequestMultiData `json:"response"`
	FeeProgramIndicator      string                                 `json:"feeProgramIndicator,omitempty"` // Optional: string | Length:3
}

type PullFundsTransactionRequestMultiData struct {
	LocalTransactionDateTime      string                    `json:"localTransactionDateTime"`                // RFC3339. dateTime | YYYY-MM-DDThh:mm:ss. The date and time you submit the transaction
	SystemsTraceAuditNumber       int                       `json:"systemsTraceAuditNumber"`                 // required, 6
	RetrievalReferenceNumber      string                    `json:"retrievalReferenceNumber"`                // ydddhhnnnnnn(numeric characters only), Length: 12
	SenderPrimaryAccountNumber    string                    `json:"senderPrimaryAccountNumber"`              // string | Length: 13 - 19
	SenderCardExpiryDate          string                    `json:"senderCardExpiryDate"`                    // string | YYYY-MM
	SenderCurrencyCode            string                    `json:"senderCurrencyCode"`                      // string | Length: 3
	Amount                        float64                   `json:"amount,omitempty"`                        // Optional: decimal | Length: totalDigits 12, fractionDigits 3 (minimum value is 0)
	Surcharge                     float64                   `json:"surcharge,omitempty"`                     // Optional: decimal | Length: totalDigits 12, fractionDigits 3(minimum value is 0)
	Cavv                          string                    `json:"cavv"`                                    // string | Length:40
	ForeignExchangeFeeTransaction float64                   `json:"foreignExchangeFeeTransaction,omitempty"` // Optional: decimal | Length: totalDigits 12, fractionDigits 3 (minimum value is 0)
	CardAcceptor                  CardAcceptor              `json:"cardAcceptor"`                            // Object
	MagneticStripeData            *MagneticStripeData       `json:"magneticStripeData,omitempty"`            // Optional: Object
	PointOfServiceData            *PointOfServiceData       `json:"pointOfServiceData,omitempty"`            // Conditional: Object
	PointOfServiceCapability      *PointOfServiceCapability `json:"pointOfServiceCapability,omitempty"`      // Conditional: Object
	PinData                       *PinData                  `json:"pinData,omitempty"`                       // Conditional: Object
}

type PushFundsTransactionRequestMultiData struct {
	SystemsTraceAuditNumber       int                       `json:"systemsTraceAuditNumber"`            // required, 6
	RetrievalReferenceNumber      string                    `json:"retrievalReferenceNumber"`           // ydddhhnnnnnn(numeric characters only), Length: 12
	LocalTransactionDateTime      string                    `json:"localTransactionDateTime"`           // RFC3339. dateTime | YYYY-MM-DDThh:mm:ss. The date and time you submit the transaction
	SenderAccountNumber           string                    `json:"senderAccountNumber,omitempty"`      // Conditional: string | Length: 0 - 34
	SenderAddress                 string                    `json:"senderAddress,omitempty"`            // Conditional: string | Length: 1 to 35
	SenderCity                    string                    `json:"senderCity,omitempty"`               // Conditional: string | Length: 1 to 25
	SenderStateCode               string                    `json:"senderStateCode,omitempty"`          // Optional: string | Length: 2
	SenderCountryCode             string                    `json:"senderCountryCode,omitempty"`        // Optional: string | Length: 2 or 33
	SenderName                    string                    `json:"senderName,omitempty"`               // Optional: string | Length: 1 to 30
	SenderReference               string                    `json:"senderReference,omitempty"`          // Optional: string | only alphabets (a-z, A-Z) and/or numbers (0-9) allowed , max: 16 characters
	SenderDateOfBirth             string                    `json:"senderDateOfBirth,omitempty"`        // Optional: string | YYYY-MM-DD
	RecipientName                 string                    `json:"recipientName,omitempty"`            // Conditional: string | Length: minimum 1, maximum 30
	RecipientPrimaryAccountNumber string                    `json:"recipientPrimaryAccountNumber"`      // string | Length: 13 - 19
	TransactionIdentifier         int                       `json:"transactionIdentifier"`              // integer | positive, Length: 15
	TransactionCurrencyCode       string                    `json:"transactionCurrencyCode"`            // string | Length: 3
	SourceOfFundsCode             string                    `json:"sourceOfFundsCode,omitempty"`        // Conditional: string | Length: 2
	Amount                        float64                   `json:"amount,omitempty"`                   // Optional: decimal | Length: totalDigits 12, fractionDigits 3 (minimum value is 0)
	CardAcceptor                  CardAcceptor              `json:"cardAcceptor"`                       // Object
	MagneticStripeData            *MagneticStripeData       `json:"magneticStripeData,omitempty"`       // Optional: Object
	PointOfServiceData            *PointOfServiceData       `json:"pointOfServiceData,omitempty"`       // Conditional: Object
	PointOfServiceCapability      *PointOfServiceCapability `json:"pointOfServiceCapability,omitempty"` // Conditional: Object
	PinData                       *PinData                  `json:"pinData,omitempty"`                  // Conditional: Object
	FeeProgramIndicator           string                    `json:"feeProgramIndicator,omitempty"`      // Optional: string | Length:3
}

type CardAcceptor struct {
	Name       string              `json:"name"`       // string | Length: 1 - 25
	TerminalId string              `json:"terminalId"` // string | Length: 1 - 8
	IdCode     string              `json:"idCode"`     // string | Length: 1 - 15
	Address    CardAcceptorAddress `json:"address"`    // Object
}

type CardAcceptorAddress struct {
	State   string `json:"state,omitempty"`   // Conditional: string | Length: 2
	County  string `json:"county,omitempty"`  // Conditional: string | Length: 3
	Country string `json:"country"`           // string | Length: 3
	ZipCode string `json:"zipCode,omitempty"` // Conditional: string | Length: 5 - 9
}

type MagneticStripeData struct {
	Track1Data string `json:"track1Data,omitempty"` // Conditional: string | Length: maximum 76
	Track2Data string `json:"track2Data,omitempty"` // Conditional: string | hex binary value is sent as String, Length: maximum 19
}

type PointOfServiceData struct {
	PanEntryMode     int    `json:"panEntryMode,omitempty"`     // Conditional: integer | positive, Length: totaldigits 2
	PosConditionCode int    `json:"posConditionCode,omitempty"` // Conditional: integer | positive,Length: totalDigits 2
	MotoECIIndicator string `json:"motoECIIndicator,omitempty"` // Conditional: string | Length: 1 , max: 1 characters
}

type PointOfServiceCapability struct {
	PosTerminalType            int `json:"posTerminalType,omitempty"`            // Conditional: integer | positive, totalDigits 0
	PosTerminalEntryCapability int `json:"posTerminalEntryCapability,omitempty"` // Conditional: integer | positive, Length: totalDigits 1
}

type PinData struct {
	PinDataBlock               string                      `json:"pinDataBlock,omitempty"`               // Conditional: string | Length: 16, hexbinary format
	SecurityRelatedControlInfo *SecurityRelatedControlInfo `json:"securityRelatedControlInfo,omitempty"` // Conditional: object
}

type SecurityRelatedControlInfo struct {
	PinBlockFormatCode int `json:"pinBlockFormatCode,omitempty"` // Conditional: integer |positive Length: totalDigits 2
	ZoneKeyIndex       int `json:"zoneKeyIndex,omitempty"`       // Conditional: integer |positive Length: totalDigits 2
}

type PullFundsTransactionResponse struct {
	StatusIdentifier      string `json:"statusIdentifier"`              // string | required when call times out
	TransactionIdentifier int    `json:"transactionIdentifier"`         // integer | positive and required when call does not timeout, Length: 15
	ActionCode            string `json:"actionCode"`                    // string | Length: 2
	ApprovalCode          string `json:"ApprovalCode,omitempty"`        // Optional: string | Length: 6
	TransmissionDateTime  string `json:"transmissionDateTime"`          // dateTime | YYYY-MM-DDThh:mm:ss
	CavvResultCode        string `json:"cavvResultCode,omitempty"`      // Optional: string | Length: 1
	ResponseCode          string `json:"responseCode"`                  // string | Length: 1
	FeeProgramIndicator   string `json:"feeProgramIndicator,omitempty"` // Optional: string | Length:3
	ErrorMessage          string `json:"errorMessage,omitempty"`        // Optional: string | Length:3
}

type PushFundsTransactionResponse struct {
	StatusIdentifier       string `json:"statusIdentifier"`              // string | required when call times out
	TransactionIdentifier  int    `json:"transactionIdentifier"`         // integer | positive and required when call does not timeout, Length: 15
	ActionCode             string `json:"actionCode"`                    // string | Length: 2
	ApprovalCode           string `json:"ApprovalCode,omitempty"`        // Optional: string | Length: 6
	TransmissionDateTime   string `json:"transmissionDateTime"`          // dateTime | YYYY-MM-DDThh:mm:ss
	CavvResultCode         string `json:"cavvResultCode,omitempty"`      // Optional: string | Length: 1
	ResponseCode           string `json:"responseCode"`                  // string | Length: 1
	FeeProgramIndicator    string `json:"feeProgramIndicator,omitempty"` // Optional: string | Length:3
	ErrorMessage           string `json:"errorMessage,omitempty"`        // Optional: string | Length:3
	PrepaidBalanceCurrency string `json:"prepaidBalanceCurrency"`        // string | Length:3
	PrepaidBalance         string `json:"prepaidBalance"`                // string
}

type ReverseFundsTransactionResponse struct {
	StatusIdentifier      string `json:"statusIdentifier"`              // string | required when call times out
	TransactionIdentifier int    `json:"transactionIdentifier"`         // integer | positive and required when call does not timeout, Length: 15
	ActionCode            string `json:"actionCode"`                    // string | Length: 2
	ApprovalCode          string `json:"ApprovalCode,omitempty"`        // Optional: string | Length: 6
	TransmissionDateTime  string `json:"transmissionDateTime"`          // dateTime | YYYY-MM-DDThh:mm:ss
	ResponseCode          string `json:"responseCode"`                  // string | Length: 1
	FeeProgramIndicator   string `json:"feeProgramIndicator,omitempty"` // Optional: string | Length:3
	ErrorMessage          string `json:"errorMessage,omitempty"`        // Optional: string | Length:3
}

// PullFundsTransactions (POST) Resource debits (pulls) funds from a sender's Visa account (in preparation for pushing funds to a recipient's account)
// by initiating a financial message called an Account Funding Transaction (AFT)
func PullFundsTransactionsPost(request PullFundsTransactionRequest) (response PullFundsTransactionResponse, err error) {
	/*
	   You should log or otherwise retain all the information returned in the PullFundsTransactions response.
	   Should it be necessary to initiate a ReverseFundsTransactions POST operation, you may need to provide
	   several of the PullFundsTransactions Response values in the Request.
	*/
	body, err := json.Marshal(request)
	if err != nil {
		return response, err
	}
	responseJson, err := Client(USER_ID, USER_PASSWORD, PULL_FUNDS_TRANSACTIONS_URL, "POST", false, body, "0")
	if err != nil {
		return response, err
	}
	// Unmarshall response
	err = json.Unmarshal(responseJson, &response)
	if err != nil {
		return response, err
	}
	return
}

// PullFundsTransactions (GET) gets the status and details for a specific PullFundsTransactions request
// if there was a timeout or connection issue on the initial request
func PullFundsTransactionsGet(statusIdentifier string) (response PullFundsTransactionResponse, err error) {
	/*
	   Get the status and details for a specific PullFundsTransactions POST request if there was a timeout or connection issue on the initial request
	   The PullFundsTransactions GET operation can be invoked when the initial PullFundsTransactions POST request has returned a timeout error.
	   When a timeout occurs, the response will include the appropriate PullFundsTransactions link header which the client then uses to get the
	   status and details of the initial request.
	*/
	requestUrl := PULL_FUNDS_TRANSACTIONS_URL + statusIdentifier
	responseJson, err := Client(USER_ID, USER_PASSWORD, requestUrl, "GET", false, nil, "0")
	if err != nil {
		return response, err
	}
	// Unmarshall response
	err = json.Unmarshal(responseJson, &response)
	if err != nil {
		return response, err
	}
	return
}

//The MultiPullFundsTransactions resource debits (pulls) funds from multiple sender's Visa accounts (in preparation for pushing funds to one or many
//recipient's accounts) by initiating an extension of the Account Funding Transaction (AFT) financial message. The MultiPullFundsTransactions
//resource can be used to submit large API requests with multiple transactions to gain operational efficiencies.
func MultiPullFundsTransactionsPost(request PullFundsTransactionRequestMulti) (response PullFundsTransactionResponse, err error) {
	/*
	   Same functionality as PullFundsTransactionsPost but with multiple requests
	*/
	body, err := json.Marshal(request)
	if err != nil {
		return response, err
	}
	responseJson, err := Client(USER_ID, USER_PASSWORD, PULL_MULTI_FUNDS_TRANSACTIONS_URL, "POST", false, body, "0")
	if err != nil {
		return response, err
	}
	// Unmarshall response
	err = json.Unmarshal(responseJson, &response)
	if err != nil {
		return response, err
	}
	return
}

//The MultiPullFundsTransactionsGet gets the status and details for a specific MultiPullFundsTransactions POST request.
func MultiPullFundsTransactionsGet(statusIdentifier string) (response PullFundsTransactionRequestMultiResponse, err error) {
	requestUrl := PULL_MULTI_FUNDS_TRANSACTIONS_URL + statusIdentifier
	responseJson, err := Client(USER_ID, USER_PASSWORD, requestUrl, "GET", false, nil, "0")
	if err != nil {
		return response, err
	}
	// Unmarshall response
	err = json.Unmarshal(responseJson, &response)
	if err != nil {
		return response, err
	}
	return
}

//PushFundsTransactions resource credits (pushes) funds to a recipient's Visa account by
//initiating a financial message called an Original Credit Transaction (OCT)
func PushFundsTransactionsPost(request PushFundsTransactionRequest) (response PushFundsTransactionResponse, err error) {
	body, err := json.Marshal(request)
	if err != nil {
		return response, err
	}
	responseJson, err := Client(USER_ID, USER_PASSWORD, PUSH_FUNDS_TRANSACTIONS_URL, "POST", false, body, "0")
	if err != nil {
		return response, err
	}
	// Unmarshall response
	err = json.Unmarshal(responseJson, &response)
	if err != nil {
		return response, err
	}
	return
}

func PushFundsTransactionsGet(statusIdentifier string) (response PushFundsTransactionResponse, err error) {
	requestUrl := PUSH_FUNDS_TRANSACTIONS_URL + statusIdentifier
	responseJson, err := Client(USER_ID, USER_PASSWORD, requestUrl, "GET", false, nil, "0")
	if err != nil {
		return response, err
	}
	// Unmarshall response
	err = json.Unmarshal(responseJson, &response)
	if err != nil {
		return response, err
	}
	return
}

//The MultiPushFundsTransactions resource credits (pushes) funds to multiple recipient's Visa accounts by initiating an extension of the
//Original Credit Transaction (OCT) financial message. The MultiPushFundsTransactions resource can be used to submit large API
//requests with multiple transactions to gain operational efficiencies.
func MultiPushFundsTransactionsPost(request PushFundsTransactionRequestMulti) (response PushFundsTransactionResponse, err error) {
	body, err := json.Marshal(request)
	if err != nil {
		return response, err
	}
	responseJson, err := Client(USER_ID, USER_PASSWORD, PUSH_MULTI_FUNDS_TRANSACTIONS_URL, "POST", false, body, "0")
	if err != nil {
		return response, err
	}
	// Unmarshall response
	err = json.Unmarshal(responseJson, &response)
	if err != nil {
		return response, err
	}
	return
}

//The MultiPushFundsTransactions GET operation can be invoked when the initial MultiPushFundsTransactions POST request has been
//submitted successfully. All successful MultiPushFundsTransactions POST requests will return a simple 202 response with the
//appropriate MultiPushFundsTransactions Link header which the client then uses it to get the status and details of the initial reques
func MultiPushFundsTransactionsGet(statusIdentifier string) (response PushFundsTransactionRequestMultiResponse, err error) {
	requestUrl := PUSH_MULTI_FUNDS_TRANSACTIONS_URL + statusIdentifier
	responseJson, err := Client(USER_ID, USER_PASSWORD, requestUrl, "GET", false, nil, "0")
	if err != nil {
		return response, err
	}
	// Unmarshall response
	err = json.Unmarshal(responseJson, &response)
	if err != nil {
		return response, err
	}
	return
}

//The ReverseFundsTransactions resource credits (pushes back) funds to the sender's Visa account by initiating a
//financial message called an Account Funding Transaction Reversal (AFTR).
func ReverseFundsTransactionsPost(request ReverseFundsTransactionRequest) (response ReverseFundsTransactionResponse, err error) {
	body, err := json.Marshal(request)
	if err != nil {
		return response, err
	}
	responseJson, err := Client(USER_ID, USER_PASSWORD, REVERSE_FUNDS_TRANSACTIONS_URL, "POST", false, body, "0")
	if err != nil {
		return response, err
	}
	// Unmarshall response
	err = json.Unmarshal(responseJson, &response)
	if err != nil {
		return response, err
	}
	return
}

//Get the status and details for a specific ReverseFundsTransactions POST request.
//The ReverseFundsTransactions GET operation can be invoked when the initial ReverseFundsTransactions POST request has returned a
//timeout error. When a timeout occurs, the response will include the appropriate reverseFundsTransactions Link header which the client
//then uses to get the status and details of the initial request.
func ReverseFundsTransactionsGet(statusIdentifier string) (response ReverseFundsTransactionResponse, err error) {
	requestUrl := REVERSE_FUNDS_TRANSACTIONS_URL + statusIdentifier
	responseJson, err := Client(USER_ID, USER_PASSWORD, requestUrl, "GET", false, nil, "0")
	if err != nil {
		return response, err
	}
	// Unmarshall response
	err = json.Unmarshal(responseJson, &response)
	if err != nil {
		return response, err
	}
	return
}
