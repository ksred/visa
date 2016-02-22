package visa

import (
	"encoding/json"
	"fmt"
	"log"
)

var PULL_FUNDS_TRANSACTIONS_URL = API_URL + "/visadirect/fundstransfer/v1/pullfundstransactions/"

type FundsTransactionRequest struct {
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

type FundsTransactionResponse struct {
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

// PullFundsTransactions (POST) Resource debits (pulls) funds from a sender's Visa account (in preparation for pushing funds to a recipient's account)
// by initiating a financial message called an Account Funding Transaction (AFT)
func PullFundsTransactionsPost(request FundsTransactionRequest) (response FundsTransactionResponse) {
	/*
	   You should log or otherwise retain all the information returned in the PullFundsTransactions response.
	   Should it be necessary to initiate a ReverseFundsTransactions POST operation, you may need to provide
	   several of the PullFundsTransactions Response values in the Request.
	*/
	body, err := json.Marshal(request)
	if err != nil {
		// @TODO: No not fatal log
		log.Fatalf("Error json encoding PullFundsTransactionsPost request: %v", err)
		return
	}
	responseJson := Client(USER_ID, USER_PASSWORD, PULL_FUNDS_TRANSACTIONS_URL, "POST", false, body, "0")
	// Unmarshall response
	err = json.Unmarshal(responseJson, &response)
	if err != nil {
		log.Fatalf("Error json decoding PullFundsTransactionsPost response: %v", err)
		return
	}
	return
}

// PullFundsTransactions (GET) gets the status and details for a specific PullFundsTransactions request
// if there was a timeout or connection issue on the initial request
func PullFundsTransactionsGet(statusIdentifier string) (response FundsTransactionResponse) {
	/*
	   Get the status and details for a specific PullFundsTransactions POST request if there was a timeout or connection issue on the initial request
	   The PullFundsTransactions GET operation can be invoked when the initial PullFundsTransactions POST request has returned a timeout error.
	   When a timeout occurs, the response will include the appropriate PullFundsTransactions link header which the client then uses to get the
	   status and details of the initial request.
	*/
	requestUrl := PULL_FUNDS_TRANSACTIONS_URL + statusIdentifier
	fmt.Printf("URL: %s\n", requestUrl)
	responseJson := Client(USER_ID, USER_PASSWORD, requestUrl, "GET", false, nil, "0")
	// Unmarshall response
	err := json.Unmarshal(responseJson, &response)
	if err != nil {
		log.Fatalf("Error json decoding PullFundsTransactionsGet response: %v", err)
		return
	}
	return
}
