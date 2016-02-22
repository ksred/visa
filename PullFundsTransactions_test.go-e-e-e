package visa

import (
	"fmt"
	"reflect"
	"testing"
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
		{},
		{
			123456,
			"123456789012",
			"2016-02-12T16:22:13",
			501111,
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
		request := FundsTransactionRequest{
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

		setVariables("USER_ID", "USER_PASSWORD")
		response := PullFundsTransactionsPost(request)
		fmt.Println(response)
		// 1. Check type
		if reflect.TypeOf(response).String() != "visa.FundsTransactionResponse" {
			t.Errorf("Return should be of type FundsTransactionResponse. Looking for %s, got %s", "visa.FundsTransactionResponse", reflect.TypeOf(response).String())
		}
	}
}
