package visa

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestPullFundsTransactionPost(t *testing.T) {
	cases := []struct {
		systemsTraceAuditNumber        int
		retrievalReferenceNumber       string
		localTransactionDateTime       string
		acquiringBin                   int
		acquirerCountryCode            int
		senderPrimaryAccountNumber     string
		senderCardExpiryDate           string
		senderCurrencyCode             string
		amount                         float64
		surcharge                      float64
		cavv                           string
		foreignExchangeFeeTransaction  float64
		businessApplicationId          string
		merchantCategoryCode           int
		CAname                         string
		CAterminalId                   string
		CAidCode                       string
		CAAstate                       string
		CAAcounty                      string
		CAAcountry                     string
		CAAzipCode                     string
		MSDtrack1Data                  string
		MSDtrack2Data                  string
		POSDpanEntryMode               int
		POSDposConditionCode           int
		POSDmotoECIIndicator           string
		POSCposTerminalType            int
		POSCposTerminalEntryCapability int
		PDpinDataBlock                 string
		PDSRCIpinBlockFormatCode       int
		PDSRCIzoneKeyIndex             int
		feeProgramIndicator            string
	}{
		//{},
		{
			123456,
			"407509300259",
			time.Now().Format("2006-01-02T03:04:05"),
			409999,
			101,
			"4957030100009952",
			"2020-03",
			"USD",
			110.,
			2.00,
			"0000010926000071934977253000000000000000",
			10.00,
			"AA",
			6012,
			"Saranya",
			"365539",
			"VMT200911026070",
			"CA",
			"081",
			"USA",
			"94404",
			"", //"1010101010101010101010101010",
			"",
			90,
			0,
			"0",
			4,
			2,
			"",
			0,
			0,
			"123",
		},
	}

	for _, c := range cases {
		cardAcceptorAddress := CardAcceptorAddress{
			State:   c.CAAstate,
			County:  c.CAAcounty,
			Country: c.CAAcountry,
			ZipCode: c.CAAzipCode,
		}
		cardAcceptor := CardAcceptor{
			Name:       c.CAname,
			TerminalId: c.CAterminalId,
			IdCode:     c.CAidCode,
			Address:    cardAcceptorAddress,
		}
		/*magneticStripeData := MagneticStripeData{
			Track1Data: c.MSDtrack1Data,
			Track2Data: c.MSDtrack2Data,
		}*/
		/*
			pointOfServiceData := PointOfServiceData{
				PanEntryMode:     c.POSDpanEntryMode,
				PosConditionCode: c.POSDposConditionCode,
				MotoECIIndicator: c.POSDmotoECIIndicator,
			}
		*/
		/*
			pointOfServiceCapability := PointOfServiceCapability{
				PosTerminalType:            c.POSCposTerminalType,
				PosTerminalEntryCapability: c.POSCposTerminalEntryCapability,
			}
		*/
		/*
			securityRelatedControlInfo := SecurityRelatedControlInfo{
				PinBlockFormatCode: c.PDSRCIpinBlockFormatCode,
				ZoneKeyIndex:       c.PDSRCIzoneKeyIndex,
			}
		*/
		/*
			pinData := PinData{
				PinDataBlock: c.PDpinDataBlock,
				//SecurityRelatedControlInfo: securityRelatedControlInfo,
			}
		*/
		request := PullFundsTransactionRequest{
			SystemsTraceAuditNumber:    c.systemsTraceAuditNumber,
			RetrievalReferenceNumber:   c.retrievalReferenceNumber,
			LocalTransactionDateTime:   c.localTransactionDateTime,
			AcquiringBin:               c.acquiringBin,
			AcquirerCountryCode:        c.acquirerCountryCode,
			SenderPrimaryAccountNumber: c.senderPrimaryAccountNumber,
			SenderCardExpiryDate:       c.senderCardExpiryDate,
			SenderCurrencyCode:         c.senderCurrencyCode,
			Amount:                     c.amount,
			Surcharge:                  c.surcharge,
			Cavv:                       c.cavv,
			ForeignExchangeFeeTransaction: c.foreignExchangeFeeTransaction,
			BusinessApplicationId:         c.businessApplicationId,
			MerchantCategoryCode:          c.merchantCategoryCode,
			CardAcceptor:                  cardAcceptor,
			//MagneticStripeData:            &magneticStripeData,
			//PointOfServiceData:       &pointOfServiceData,
			//PointOfServiceCapability: &pointOfServiceCapability,
			//PinData:                  &pinData,
			FeeProgramIndicator: c.feeProgramIndicator,
		}

		setVariables(TEST_USER_KEY, TEST_USER_PASSWORD)
		response, err := PullFundsTransactionsPost(request)
		if err != nil {
			t.Errorf("Error when getting response: %v", err)
		}
		fmt.Printf("%+v\n", response)
		// 1. Check type
		if reflect.TypeOf(response).String() != "visa.PullFundsTransactionResponse" {
			t.Errorf("Return should be of type PullFundsTransactionResponse. Looking for %s, got %s", "visa.PullFundsTransactionResponse", reflect.TypeOf(response).String())
		}
	}
}

//@FIXME: Currently receiving a 404. Suspect this might have to do with the transaction not being found, as opposed to the path
func TestPullFundsTransactionGet(t *testing.T) {
	cases := []struct {
		statusIdentifier string
	}{
		{
			"381228649430011",
		},
	}

	for _, c := range cases {
		setVariables(TEST_USER_KEY, TEST_USER_PASSWORD)
		response, err := PullFundsTransactionsGet(c.statusIdentifier)
		if err != nil {
			t.Errorf("Error when getting response: %v", err)
		}
		fmt.Printf("%+v\n", response)
		// 1. Check type
		if reflect.TypeOf(response).String() != "visa.PullFundsTransactionResponse" {
			t.Errorf("Return should be of type PullFundsTransactionResponse. Looking for %s, got %s", "visa.PullFundsTransactionResponse", reflect.TypeOf(response).String())
		}
	}
}

func TestPullFundsTransactionMultiPost(t *testing.T) {
	cases := []struct {
		systemsTraceAuditNumber        int
		retrievalReferenceNumber       string
		localTransactionDateTime       string
		acquiringBin                   int
		acquirerCountryCode            int
		senderPrimaryAccountNumber     string
		senderCardExpiryDate           string
		senderCurrencyCode             string
		amount                         float64
		surcharge                      float64
		cavv                           string
		foreignExchangeFeeTransaction  float64
		businessApplicationId          string
		merchantCategoryCode           int
		CAname                         string
		CAterminalId                   string
		CAidCode                       string
		CAAstate                       string
		CAAcounty                      string
		CAAcountry                     string
		CAAzipCode                     string
		MSDtrack1Data                  string
		MSDtrack2Data                  string
		POSDpanEntryMode               int
		POSDposConditionCode           int
		POSDmotoECIIndicator           string
		POSCposTerminalType            int
		POSCposTerminalEntryCapability int
		PDpinDataBlock                 string
		PDSRCIpinBlockFormatCode       int
		PDSRCIzoneKeyIndex             int
		feeProgramIndicator            string
	}{
		//{},
		{
			123456,
			"407509300259",
			time.Now().Format("2006-01-02T03:04:05"),
			409999,
			101,
			"4957030100009952",
			"2020-03",
			"USD",
			110.,
			2.00,
			"0000010926000071934977253000000000000000",
			10.00,
			"AA",
			6012,
			"Saranya",
			"365539",
			"VMT200911026070",
			"CA",
			"081",
			"USA",
			"94404",
			"", //"1010101010101010101010101010",
			"",
			90,
			0,
			"0",
			4,
			2,
			"",
			0,
			0,
			"123",
		},
	}

	for _, c := range cases {
		cardAcceptorAddress := CardAcceptorAddress{
			State:   c.CAAstate,
			County:  c.CAAcounty,
			Country: c.CAAcountry,
			ZipCode: c.CAAzipCode,
		}
		cardAcceptor := CardAcceptor{
			Name:       c.CAname,
			TerminalId: c.CAterminalId,
			IdCode:     c.CAidCode,
			Address:    cardAcceptorAddress,
		}
		/*magneticStripeData := MagneticStripeData{
			Track1Data: c.MSDtrack1Data,
			Track2Data: c.MSDtrack2Data,
		}*/
		/*
			pointOfServiceData := PointOfServiceData{
				PanEntryMode:     c.POSDpanEntryMode,
				PosConditionCode: c.POSDposConditionCode,
				MotoECIIndicator: c.POSDmotoECIIndicator,
			}
		*/
		/*
			pointOfServiceCapability := PointOfServiceCapability{
				PosTerminalType:            c.POSCposTerminalType,
				PosTerminalEntryCapability: c.POSCposTerminalEntryCapability,
			}
		*/
		/*
			securityRelatedControlInfo := SecurityRelatedControlInfo{
				PinBlockFormatCode: c.PDSRCIpinBlockFormatCode,
				ZoneKeyIndex:       c.PDSRCIzoneKeyIndex,
			}
		*/
		/*
			pinData := PinData{
				PinDataBlock: c.PDpinDataBlock,
				//SecurityRelatedControlInfo: securityRelatedControlInfo,
			}
		*/
		requestData1 := PullFundsTransactionRequestMultiData{
			LocalTransactionDateTime:   c.localTransactionDateTime,
			SystemsTraceAuditNumber:    c.systemsTraceAuditNumber,
			RetrievalReferenceNumber:   c.retrievalReferenceNumber,
			SenderPrimaryAccountNumber: c.senderPrimaryAccountNumber,
			SenderCardExpiryDate:       c.senderCardExpiryDate,
			SenderCurrencyCode:         c.senderCurrencyCode,
			Amount:                     c.amount,
			//Surcharge:                  c.surcharge,
			Cavv: c.cavv,
			//ForeignExchangeFeeTransaction: c.foreignExchangeFeeTransaction,
			CardAcceptor: cardAcceptor,
			//MagneticStripeData:            &magneticStripeData,
			//PointOfServiceData:       &pointOfServiceData,
			//PointOfServiceCapability: &pointOfServiceCapability,
			//PinData:                  &pinData,
		}
		requestData2 := PullFundsTransactionRequestMultiData{
			LocalTransactionDateTime:   c.localTransactionDateTime,
			SystemsTraceAuditNumber:    c.systemsTraceAuditNumber,
			RetrievalReferenceNumber:   c.retrievalReferenceNumber,
			SenderPrimaryAccountNumber: c.senderPrimaryAccountNumber,
			SenderCardExpiryDate:       c.senderCardExpiryDate,
			SenderCurrencyCode:         c.senderCurrencyCode,
			Amount:                     c.amount,
			//Surcharge:                  c.surcharge,
			Cavv: c.cavv,
			//ForeignExchangeFeeTransaction: c.foreignExchangeFeeTransaction,
			CardAcceptor: cardAcceptor,
			//MagneticStripeData:            &magneticStripeData,
			//PointOfServiceData:       &pointOfServiceData,
			//PointOfServiceCapability: &pointOfServiceCapability,
			//PinData:                  &pinData,
		}
		request := PullFundsTransactionRequestMulti{
			LocalTransactionDateTime: c.localTransactionDateTime,
			AcquiringBin:             c.acquiringBin,
			AcquirerCountryCode:      c.acquirerCountryCode,
			BusinessApplicationId:    c.businessApplicationId,
			MerchantCategoryCode:     c.merchantCategoryCode,
			//FeeProgramIndicator:      c.feeProgramIndicator,
			Request: []PullFundsTransactionRequestMultiData{requestData1, requestData2},
		}

		type PullFundsTransactionRequestMulti struct {
			SystemsTraceAuditNumber  int                                    `json:"systemsTraceAuditNumber"`       // required, 6
			RetrievalReferenceNumber string                                 `json:"retrievalReferenceNumber"`      // ydddhhnnnnnn(numeric characters only), Length: 12
			LocalTransactionDateTime string                                 `json:"localTransactionDateTime"`      // RFC3339. dateTime | YYYY-MM-DDThh:mm:ss. The date and time you submit the transaction
			AcquiringBin             int                                    `json:"acquiringBin"`                  // integer | positive, Length: 6 - 11
			AcquirerCountryCode      int                                    `json:"acquirerCountryCode"`           // integer | Length: 3
			FeeProgramIndicator      string                                 `json:"feeProgramIndicator,omitempty"` // Optional: string | Length:3
			Request                  []PullFundsTransactionRequestMultiData `json:"request"`
		}

		type PullFundsTransactionRequestMultiData struct {
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
		}

		setVariables(TEST_USER_KEY, TEST_USER_PASSWORD)
		response, err := MultiPullFundsTransactionsPost(request)
		if err != nil {
			t.Errorf("Error when getting response: %v\n", err)
		}
		fmt.Printf("%+v\n", response)
		// 1. Check type
		if reflect.TypeOf(response).String() != "visa.PullFundsTransactionResponse" {
			t.Errorf("Return should be of type PullFundsTransactionResponse. Looking for %s, got %s", "visa.PullFundsTransactionResponse", reflect.TypeOf(response).String())
		}
	}
}

//@FIXME: Currently receiving a 404. Suspect this might have to do with the transaction not being found, as opposed to the path
func TestPullMultiFundsTransactionGet(t *testing.T) {
	cases := []struct {
		statusIdentifier string
	}{
		{
			"381228649430011",
		},
	}

	for _, c := range cases {
		setVariables(TEST_USER_KEY, TEST_USER_PASSWORD)
		response, err := MultiPullFundsTransactionsGet(c.statusIdentifier)
		if err != nil {
			t.Errorf("Error when getting response: %v", err)
		}
		fmt.Printf("%+v\n", response)
		// 1. Check type
		if reflect.TypeOf(response).String() != "visa.PullFundsTransactionRequestMultiResponse" {
			t.Errorf("Return should be of type PullFundsTransactionRequestMultiResponse. Looking for %s, got %s", "visa.PullFundsTransactionRequestMultiResponse", reflect.TypeOf(response).String())
		}
	}
}
