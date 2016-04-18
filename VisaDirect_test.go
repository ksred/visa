package visa

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/satori/go.uuid"
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
			451000,
			"330000550000",
			time.Now().Format("2006-01-02T03:04:05"),
			408999,
			840,
			"4005520000011126",
			"2020-03",
			"USD",
			100.,
			0.00,
			"0000010926000071934977253000000000000000",
			0.00,
			"AA",
			6012,
			"Acceptor 1",
			"365539",
			"VMT200911026070",
			"CA",
			"081",
			"USA",
			"94404",
			"1010101010101010101010101010", //"1010101010101010101010101010",
			"4008310000000007D130310191014085",
			90,
			0,
			"0",
			4,
			2,
			"1cd948f2b961b682",
			1,
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

		// Set UUID
		newUuid := uuid.NewV4()
		uuid := newUuid.String()

		setVariables(TEST_USER_KEY, TEST_USER_PASSWORD)
		response, err := PullFundsTransactionsPost(request, uuid)
		if err != nil {
			t.Errorf("Error when getting response: %v", err)
		}
		//fmt.Printf("%+v\n", response)
		// 1. Check type
		if reflect.TypeOf(response).String() != "visa.PullFundsTransactionResponse" {
			t.Errorf("Return should be of type PullFundsTransactionResponse. Looking for %s, got %s", "visa.PullFundsTransactionResponse", reflect.TypeOf(response).String())
		}
	}
}

func TestPullFundsTransactionGet(t *testing.T) {
	cases := []struct {
		statusIdentifier string
	}{
		{
			"234234322342343",
		},
	}

	for _, c := range cases {
		setVariables(TEST_USER_KEY, TEST_USER_PASSWORD)
		response, err := PullFundsTransactionsGet(c.statusIdentifier)
		if err != nil {
			t.Errorf("Error when getting response: %v", err)
		}
		//fmt.Printf("%+v\n", response)
		// 1. Check type
		if reflect.TypeOf(response).String() != "visa.PullFundsTransactionResponse" {
			t.Errorf("Return should be of type PullFundsTransactionResponse. Looking for %s, got %s", "visa.PullFundsTransactionResponse", reflect.TypeOf(response).String())
		}
	}
}

func TestPullFundsTransactionMultiPost(t *testing.T) {
	t.Skip("Skipping test while waiting for issue#4 to be resolved - Internal 500 on multi requests")
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

		// Set UUID
		newUuid := uuid.NewV4()
		uuid := newUuid.String()

		setVariables(TEST_USER_KEY, TEST_USER_PASSWORD)
		//Convert to JSON for debugging
		reqJson, errJS := json.Marshal(request)
		if errJS != nil {
			//fmt.Println(errJS.Error())
		}
		fmt.Printf("%+v\n", string(reqJson))
		//os.Exit(1)
		response, err := MultiPullFundsTransactionsPost(request, uuid)
		if err != nil {
			t.Errorf("Error when getting response: %v\n", err)
		}
		//fmt.Printf("%+v\n", response)
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
		//fmt.Printf("%+v\n", response)
		// 1. Check type
		if reflect.TypeOf(response).String() != "visa.PullFundsTransactionRequestMultiResponse" {
			t.Errorf("Return should be of type PullFundsTransactionRequestMultiResponse. Looking for %s, got %s", "visa.PullFundsTransactionRequestMultiResponse", reflect.TypeOf(response).String())
		}
	}
}

func TestPushFundsTransactionPost(t *testing.T) {
	cases := []struct {
		systemsTraceAuditNumber        int
		retrievalReferenceNumber       string
		localTransactionDateTime       string
		acquiringBin                   int
		acquirerCountryCode            int
		senderAccountNumber            string
		senderAddress                  string
		senderCity                     string
		senderStateCode                string
		senderCountryCode              string
		senderName                     string
		senderReference                string
		senderDateOfBirth              string
		recipientName                  string
		recipientPrimaryAccountNumber  string
		transactionIdentifier          int
		transactionCurrencyCode        string
		sourceOfFundsCode              string
		amount                         float64
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
			451018,
			"412770451018",
			time.Now().Format("2006-01-02T03:04:05"),
			409999,
			101,
			"4957030100009952",
			"901 Metro Center Blvd",
			"Foster City",
			"CA",
			"124",
			"Mohammed Qasim",
			"",
			"", //sender dob
			"roahn",
			"4957030420210496",
			381228649430015,
			"USD",
			"05",
			110.,
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
			"",
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
		request := PushFundsTransactionRequest{
			SystemsTraceAuditNumber:       c.systemsTraceAuditNumber,
			RetrievalReferenceNumber:      c.retrievalReferenceNumber,
			LocalTransactionDateTime:      c.localTransactionDateTime,
			AcquiringBin:                  c.acquiringBin,
			AcquirerCountryCode:           c.acquirerCountryCode,
			SenderAccountNumber:           c.senderAccountNumber,
			SenderAddress:                 c.senderAddress,
			SenderCity:                    c.senderCity,
			SenderStateCode:               c.senderStateCode,
			SenderCountryCode:             c.senderCountryCode,
			SenderName:                    c.senderName,
			SenderReference:               c.senderReference,
			SenderDateOfBirth:             c.senderDateOfBirth,
			RecipientName:                 c.recipientName,
			RecipientPrimaryAccountNumber: c.recipientPrimaryAccountNumber,
			TransactionIdentifier:         c.transactionIdentifier,
			TransactionCurrencyCode:       c.transactionCurrencyCode,
			SourceOfFundsCode:             c.sourceOfFundsCode,
			Amount:                        c.amount,
			BusinessApplicationId:         c.businessApplicationId,
			MerchantCategoryCode:          c.merchantCategoryCode,
			CardAcceptor:                  cardAcceptor,
			//MagneticStripeData:            &magneticStripeData,
			//PointOfServiceData:       &pointOfServiceData,
			//PointOfServiceCapability: &pointOfServiceCapability,
			//PinData:                  &pinData,
			FeeProgramIndicator: c.feeProgramIndicator,
		}

		// Set UUID
		newUuid := uuid.NewV4()
		uuid := newUuid.String()

		setVariables(TEST_USER_KEY, TEST_USER_PASSWORD)
		response, err := PushFundsTransactionsPost(request, uuid)
		if err != nil {
			t.Errorf("Error when getting response: %v", err)
		}
		//fmt.Printf("%+v\n", response)
		// 1. Check type
		if reflect.TypeOf(response).String() != "visa.PushFundsTransactionResponse" {
			t.Errorf("Return should be of type PushFundsTransactionResponse. Looking for %s, got %s", "visa.PushFundsTransactionResponse", reflect.TypeOf(response).String())
		}
	}
}

func TestPushFundsTransactionGet(t *testing.T) {
	cases := []struct {
		statusIdentifier string
	}{
		{
			"234234322342343",
		},
	}

	for _, c := range cases {
		setVariables(TEST_USER_KEY, TEST_USER_PASSWORD)
		response, err := PushFundsTransactionsGet(c.statusIdentifier)
		if err != nil {
			t.Errorf("Error when getting response: %v", err)
		}
		//fmt.Printf("%+v\n", response)
		// 1. Check type
		if reflect.TypeOf(response).String() != "visa.PushFundsTransactionResponse" {
			t.Errorf("Return should be of type PushFundsTransactionResponse. Looking for %s, got %s", "visa.PushFundsTransactionResponse", reflect.TypeOf(response).String())
		}
	}
}

func TestPushFundsTransactionMultiPost(t *testing.T) {
	t.Skip("Skipping test while waiting for issue#4 to be resolved - Internal 500 on multi requests")
	cases := []struct {
		systemsTraceAuditNumber        int
		retrievalReferenceNumber       string
		localTransactionDateTime       string
		acquiringBin                   int
		acquirerCountryCode            int
		senderAccountNumber            string
		senderAddress                  string
		senderCity                     string
		senderStateCode                string
		senderCountryCode              string
		senderName                     string
		senderReference                string
		senderDateOfBirth              string
		recipientName                  string
		recipientPrimaryAccountNumber  string
		transactionIdentifier          int
		transactionCurrencyCode        string
		sourceOfFundsCode              string
		amount                         float64
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
			451018,
			"412770451018",
			time.Now().Format("2006-01-02T03:04:05"),
			409999,
			101,
			"4957030100009952",
			"901 Metro Center Blvd",
			"Foster City",
			"CA",
			"124",
			"Mohammed Qasim",
			"",
			"", //sender dob
			"roahn",
			"4957030420210496",
			381228649430015,
			"USD",
			"05",
			110.,
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
			"",
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
		requestData1 := PushFundsTransactionRequestMultiData{
			SystemsTraceAuditNumber:       c.systemsTraceAuditNumber,
			RetrievalReferenceNumber:      c.retrievalReferenceNumber,
			LocalTransactionDateTime:      c.localTransactionDateTime,
			SenderAccountNumber:           c.senderAccountNumber,
			SenderAddress:                 c.senderAddress,
			SenderCity:                    c.senderCity,
			SenderStateCode:               c.senderStateCode,
			SenderCountryCode:             c.senderCountryCode,
			SenderName:                    c.senderName,
			SenderReference:               c.senderReference,
			SenderDateOfBirth:             c.senderDateOfBirth,
			RecipientName:                 c.recipientName,
			RecipientPrimaryAccountNumber: c.recipientPrimaryAccountNumber,
			TransactionIdentifier:         c.transactionIdentifier,
			TransactionCurrencyCode:       c.transactionCurrencyCode,
			SourceOfFundsCode:             c.sourceOfFundsCode,
			Amount:                        c.amount,
			CardAcceptor:                  cardAcceptor,
			//MagneticStripeData:            &magneticStripeData,
			//PointOfServiceData:       &pointOfServiceData,
			//PointOfServiceCapability: &pointOfServiceCapability,
			//PinData:                  &pinData,
			FeeProgramIndicator: c.feeProgramIndicator,
		}

		requestData2 := PushFundsTransactionRequestMultiData{
			SystemsTraceAuditNumber:       c.systemsTraceAuditNumber,
			RetrievalReferenceNumber:      c.retrievalReferenceNumber,
			LocalTransactionDateTime:      c.localTransactionDateTime,
			SenderAccountNumber:           c.senderAccountNumber,
			SenderAddress:                 c.senderAddress,
			SenderCity:                    c.senderCity,
			SenderStateCode:               c.senderStateCode,
			SenderCountryCode:             c.senderCountryCode,
			SenderName:                    c.senderName,
			SenderReference:               c.senderReference,
			SenderDateOfBirth:             c.senderDateOfBirth,
			RecipientName:                 c.recipientName,
			RecipientPrimaryAccountNumber: c.recipientPrimaryAccountNumber,
			TransactionIdentifier:         c.transactionIdentifier,
			TransactionCurrencyCode:       c.transactionCurrencyCode,
			SourceOfFundsCode:             c.sourceOfFundsCode,
			Amount:                        c.amount,
			CardAcceptor:                  cardAcceptor,
			//MagneticStripeData:            &magneticStripeData,
			//PointOfServiceData:       &pointOfServiceData,
			//PointOfServiceCapability: &pointOfServiceCapability,
			//PinData:                  &pinData,
			FeeProgramIndicator: c.feeProgramIndicator,
		}

		request := PushFundsTransactionRequestMulti{
			LocalTransactionDateTime: c.localTransactionDateTime,
			AcquiringBin:             c.acquiringBin,
			AcquirerCountryCode:      c.acquirerCountryCode,
			BusinessApplicationId:    c.businessApplicationId,
			MerchantCategoryCode:     c.merchantCategoryCode,
			//FeeProgramIndicator:      c.feeProgramIndicator,
			Request: []PushFundsTransactionRequestMultiData{requestData1, requestData2},
		}

		// Set UUID
		newUuid := uuid.NewV4()
		uuid := newUuid.String()

		setVariables(TEST_USER_KEY, TEST_USER_PASSWORD)
		/*
			//Convert to JSON for debugging
			reqJson, errJS := json.Marshal(request)
			if errJS != nil {
				fmt.Println(errJS.Error())
			}
			fmt.Printf("%+v\n", string(reqJson))
			//os.Exit(1)
		*/
		response, err := MultiPushFundsTransactionsPost(request, uuid)
		if err != nil {
			t.Errorf("Error when getting response: %v", err)
		}
		//fmt.Printf("%+v\n", response)
		// 1. Check type
		if reflect.TypeOf(response).String() != "visa.PushFundsTransactionResponse" {
			t.Errorf("Return should be of type PushFundsTransactionResponse. Looking for %s, got %s", "visa.PushFundsTransactionResponse", reflect.TypeOf(response).String())
		}
	}

}

func TestPushMultiFundsTransactionGet(t *testing.T) {
	cases := []struct {
		statusIdentifier string
	}{
		{
			"381228649430011",
		},
	}

	for _, c := range cases {
		setVariables(TEST_USER_KEY, TEST_USER_PASSWORD)
		response, err := MultiPushFundsTransactionsGet(c.statusIdentifier)
		if err != nil {
			t.Errorf("Error when getting response: %v", err)
		}
		//fmt.Printf("%+v\n", response)
		// 1. Check type
		if reflect.TypeOf(response).String() != "visa.PushFundsTransactionRequestMultiResponse" {
			t.Errorf("Return should be of type PushFundsTransactionRequestMultiResponse. Looking for %s, got %s", "visa.PushFundsTransactionRequestMultiResponse", reflect.TypeOf(response).String())
		}
	}
}

func TestReverseFundsTransactionPost(t *testing.T) {
	cases := []struct {
		systemsTraceAuditNumber        int
		retrievalReferenceNumber       string
		localTransactionDateTime       string
		acquiringBin                   int
		acquirerCountryCode            int
		senderPrimaryAccountNumber     string
		senderCardExpiryDate           string
		senderCurrencyCode             string
		transactionIdentifier          int
		amount                         float64
		surcharge                      float64
		foreignExchangeFeeTransaction  float64
		odeApprovalCode                string
		odeSystemsTraceAuditNumber     int
		odeTransmissionsDateTime       string
		odeAcquiringBin                int
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
			451050,
			"330000550000",
			time.Now().Format("2006-01-02T03:04:05"),
			408999,
			608,
			"4895100000055127",
			"2015-10",
			"USD",
			381228649430011,
			110.,
			0.01,
			1.,
			"408999", //Original Data Element
			897825,
			time.Now().Format("2006-01-02T03:04:05"),
			408999,
			"Visa Inc. USA-Foster City", //CA
			"365539",
			"VMT200911026070",
			"CA",
			"San Mateo",
			"USA",
			"94404",
			"", //MSD
			"",
			0, //POS
			0,
			"",
			0,
			0,
			"", //PinBlock
			0,
			0,
			"",
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
		originalDataElements := ReverseOriginalDataElements{
			ApprovalCode:            c.odeApprovalCode,
			SystemsTraceAuditNumber: c.odeSystemsTraceAuditNumber,
			TransmissionDateTime:    c.odeTransmissionsDateTime,
			AcquiringBin:            c.odeAcquiringBin,
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
		request := ReverseFundsTransactionRequest{
			SystemsTraceAuditNumber:       c.systemsTraceAuditNumber,
			RetrievalReferenceNumber:      c.retrievalReferenceNumber,
			LocalTransactionDateTime:      c.localTransactionDateTime,
			AcquiringBin:                  c.acquiringBin,
			AcquirerCountryCode:           c.acquirerCountryCode,
			SenderPrimaryAccountNumber:    c.senderPrimaryAccountNumber,
			SenderCardExpiryDate:          c.senderCardExpiryDate,
			SenderCurrencyCode:            c.senderCurrencyCode,
			TransactionIdentifier:         c.transactionIdentifier,
			Amount:                        c.amount,
			Surcharge:                     c.surcharge,
			ForeignExchangeFeeTransaction: c.foreignExchangeFeeTransaction,
			OriginalDataElements:          originalDataElements,
			CardAcceptor:                  cardAcceptor,
			//MagneticStripeData:            &magneticStripeData,
			//PointOfServiceData:       &pointOfServiceData,
			//PointOfServiceCapability: &pointOfServiceCapability,
			//PinData:                  &pinData,
			FeeProgramIndicator: c.feeProgramIndicator,
		}

		// Set UUID
		newUuid := uuid.NewV4()
		uuid := newUuid.String()

		setVariables(TEST_USER_KEY, TEST_USER_PASSWORD)
		/*
			//Convert to JSON for debugging
			reqJson, errJS := json.Marshal(request)
			if errJS != nil {
				//fmt.Println(errJS.Error())
			}
			fmt.Printf("%+v\n", string(reqJson))
			//os.Exit(1)
		*/
		response, err := ReverseFundsTransactionsPost(request, uuid)
		if err != nil {
			t.Errorf("Error when getting response: %v", err)
		}
		//fmt.Printf("%+v\n", response)
		// 1. Check type
		if reflect.TypeOf(response).String() != "visa.ReverseFundsTransactionResponse" {
			t.Errorf("Return should be of type ReverseFundsTransactionResponse. Looking for %s, got %s", "visa.ReverseFundsTransactionResponse", reflect.TypeOf(response).String())
		}
	}
}

func TestReverseFundsTransactionGet(t *testing.T) {
	cases := []struct {
		statusIdentifier string
	}{
		{
			"234234322342343",
		},
	}

	for _, c := range cases {
		setVariables(TEST_USER_KEY, TEST_USER_PASSWORD)
		response, err := ReverseFundsTransactionsGet(c.statusIdentifier)
		if err != nil {
			t.Errorf("Error when getting response: %v", err)
		}
		//fmt.Printf("%+v\n", response)
		// 1. Check type
		if reflect.TypeOf(response).String() != "visa.ReverseFundsTransactionResponse" {
			t.Errorf("Return should be of type ReverseFundsTransactionResponse. Looking for %s, got %s", "visa.ReverseFundsTransactionResponse", reflect.TypeOf(response).String())
		}
	}
}

func TestReverseFundsTransactionMultiPost(t *testing.T) {
	t.Skip("Skipping test while waiting for issue#4 to be resolved - Internal 500 on multi requests")
	cases := []struct {
		systemsTraceAuditNumber        int
		retrievalReferenceNumber       string
		localTransactionDateTime       string
		acquiringBin                   int
		acquirerCountryCode            int
		senderPrimaryAccountNumber     string
		senderCardExpiryDate           string
		senderCurrencyCode             string
		transactionIdentifier          int
		amount                         float64
		surcharge                      float64
		foreignExchangeFeeTransaction  float64
		odeApprovalCode                string
		odeSystemsTraceAuditNumber     int
		odeTransmissionsDateTime       string
		odeAcquiringBin                int
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
			451050,
			"330000550000",
			time.Now().Format("2006-01-02T03:04:05"),
			408999,
			608,
			"4895100000055127",
			"2015-10",
			"USD",
			381228649430011,
			110.,
			0.01,
			1.,
			"408999", //Original Data Element
			897825,
			time.Now().Format("2006-01-02T03:04:05"),
			408999,
			"Visa Inc. USA-Foster City", //CA
			"365539",
			"VMT200911026070",
			"CA",
			"San Mateo",
			"USA",
			"94404",
			"", //MSD
			"",
			0, //POS
			0,
			"",
			0,
			0,
			"", //PinBlock
			0,
			0,
			"",
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
		originalDataElements := ReverseOriginalDataElements{
			ApprovalCode:            c.odeApprovalCode,
			SystemsTraceAuditNumber: c.odeSystemsTraceAuditNumber,
			TransmissionDateTime:    c.odeTransmissionsDateTime,
			AcquiringBin:            c.odeAcquiringBin,
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
		requestData1 := ReverseFundsTransactionRequestMultiData{
			LocalTransactionDateTime:      c.localTransactionDateTime,
			SystemsTraceAuditNumber:       c.systemsTraceAuditNumber,
			RetrievalReferenceNumber:      c.retrievalReferenceNumber,
			SenderPrimaryAccountNumber:    c.senderPrimaryAccountNumber,
			SenderCardExpiryDate:          c.senderCardExpiryDate,
			SenderCurrencyCode:            c.senderCurrencyCode,
			TransactionIdentifier:         c.transactionIdentifier,
			Amount:                        c.amount,
			Surcharge:                     c.surcharge,
			ForeignExchangeFeeTransaction: c.foreignExchangeFeeTransaction,
			OriginalDataElements:          originalDataElements,
			CardAcceptor:                  cardAcceptor,
			//MagneticStripeData:            &magneticStripeData,
			//PointOfServiceData:       &pointOfServiceData,
			//PointOfServiceCapability: &pointOfServiceCapability,
			//PinData:                  &pinData,
			FeeProgramIndicator: c.feeProgramIndicator,
		}

		requestData2 := ReverseFundsTransactionRequestMultiData{
			LocalTransactionDateTime:      c.localTransactionDateTime,
			SystemsTraceAuditNumber:       c.systemsTraceAuditNumber,
			RetrievalReferenceNumber:      c.retrievalReferenceNumber,
			SenderPrimaryAccountNumber:    c.senderPrimaryAccountNumber,
			SenderCardExpiryDate:          c.senderCardExpiryDate,
			SenderCurrencyCode:            c.senderCurrencyCode,
			TransactionIdentifier:         c.transactionIdentifier,
			Amount:                        c.amount,
			Surcharge:                     c.surcharge,
			ForeignExchangeFeeTransaction: c.foreignExchangeFeeTransaction,
			OriginalDataElements:          originalDataElements,
			CardAcceptor:                  cardAcceptor,
			//MagneticStripeData:            &magneticStripeData,
			//PointOfServiceData:       &pointOfServiceData,
			//PointOfServiceCapability: &pointOfServiceCapability,
			//PinData:                  &pinData,
			FeeProgramIndicator: c.feeProgramIndicator,
		}

		request := ReverseFundsTransactionRequestMulti{
			AcquiringBin:             c.acquiringBin,
			AcquirerCountryCode:      c.acquirerCountryCode,
			LocalTransactionDateTime: c.localTransactionDateTime,
			Request:                  []ReverseFundsTransactionRequestMultiData{requestData1, requestData2},
		}

		// Set UUID
		newUuid := uuid.NewV4()
		uuid := newUuid.String()

		setVariables(TEST_USER_KEY, TEST_USER_PASSWORD)
		response, err := MultiReverseFundsTransactionsPost(request, uuid)
		/* Convert to JSON for debugging
		reqJson, err := json.Marshal(request)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%+v\n", string(reqJson))
		os.Exit(1)
		*/
		if err != nil {
			t.Errorf("Error when getting response: %v", err)
		}
		//fmt.Printf("%+v\n", response)
		// 1. Check type
		if reflect.TypeOf(response).String() != "visa.ReverseFundsTransactionRequestMultiResponse" {
			t.Errorf("Return should be of type ReverseFundsTransactionRequestMultiResponse. Looking for %s, got %s", "visa.ReverseFundsTransactionRequestMultiResponse", reflect.TypeOf(response).String())
		}
	}
}

func TestReverseFundsTransactionMultiGet(t *testing.T) {
	cases := []struct {
		statusIdentifier string
	}{
		{
			"234234322342343",
		},
	}

	for _, c := range cases {
		setVariables(TEST_USER_KEY, TEST_USER_PASSWORD)
		response, err := MultiReverseFundsTransactionsGet(c.statusIdentifier)
		if err != nil {
			t.Errorf("Error when getting response: %v", err)
		}
		//fmt.Printf("%+v\n", response)
		// 1. Check type
		if reflect.TypeOf(response).String() != "visa.ReverseFundsTransactionRequestMultiResponse" {
			t.Errorf("Return should be of type ReverseFundsTransactionRequestMultiResponse. Looking for %s, got %s", "visa.ReverseFundsTransactionRequestMultiResponse", reflect.TypeOf(response).String())
		}
	}
}
