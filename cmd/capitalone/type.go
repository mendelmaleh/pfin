package main

import (
	"time"
)

type T struct {
	Entries []struct {
		AcquirerTransactionReferenceNumber     string    `json:"acquirerTransactionReferenceNumber,omitempty"`
		AuthorizationIssuerCode                string    `json:"authorizationIssuerCode,omitempty"`
		AuthorizationResponse                  string    `json:"authorizationResponse,omitempty"`
		AuthorizationStatus                    string    `json:"authorizationStatus,omitempty"`
		AuthorizationTimeStamp                 time.Time `json:"authorizationTimeStamp,omitempty"`
		AuthorizationType                      string    `json:"authorizationType,omitempty"`
		CategoryIconURL                        string    `json:"categoryIconURL,omitempty"`
		CategoryImageURL                       string    `json:"categoryImageURL,omitempty"`
		DisplayCategory                        string    `json:"displayCategory"`
		DisputedCount                          int       `json:"disputedCount"`
		HasDisputeIndicator                    bool      `json:"hasDisputeIndicator"`
		IsMemoPosted                           bool      `json:"isMemoPosted"`
		IsPrintable                            bool      `json:"isPrintable,omitempty"`
		IsReportedAsFraud                      bool      `json:"isReportedAsFraud"`
		LastIssuedCardLastFourDigits           string    `json:"lastIssuedCardLastFourDigits"`
		MaskedVirtualCardNumber                string    `json:"maskedVirtualCardNumber,omitempty"`
		StatementDescription                   string    `json:"statementDescription,omitempty"`
		TransactingCardLastFour                string    `json:"transactingCardLastFour"`
		TransactionAmount                      float64   `json:"transactionAmount"`
		TransactionBeginningStatementTimeStamp time.Time `json:"transactionBeginningStatementTimeStamp,omitempty"`
		TransactionCategoryCode                string    `json:"transactionCategoryCode,omitempty"`
		TransactionDate                        time.Time `json:"transactionDate,omitempty"`
		TransactionDebitCredit                 string    `json:"transactionDebitCredit"`
		TransactionDescription                 string    `json:"transactionDescription"`
		TransactionDisplayDate                 time.Time `json:"transactionDisplayDate"`
		TransactionLifecycleID                 string    `json:"transactionLifecycleId,omitempty"`
		TransactionMerchant                    struct {
			Address struct {
				AddressLine1 string `json:"addressLine1,omitempty"`
				City         string `json:"city"`
				CountryCode  string `json:"countryCode"`
				PostalCode   string `json:"postalCode"`
				StateCode    string `json:"stateCode"`
			} `json:"address"`
			Category         string `json:"category,omitempty"`
			CategoryCode     string `json:"categoryCode,omitempty"`
			ChainPhoneNumber string `json:"chainPhoneNumber,omitempty"`
			GeoLocation      *struct {
				Latitude  string `json:"latitude,omitempty"`
				Longitude string `json:"longitude,omitempty"`
			} `json:"geoLocation,omitempty"`
			LogoURL          string `json:"logoURL,omitempty"`
			MerchantID       string `json:"merchantId,omitempty"`
			MerchantType     string `json:"merchantType,omitempty"`
			MerchantTypeCode string `json:"merchantTypeCode,omitempty"`
			Name             string `json:"name"`
			ParentCategory   string `json:"parentCategory,omitempty"`
			PhoneNumber      string `json:"phoneNumber"`
			WebsiteURL       string `json:"websiteURL,omitempty"`
		} `json:"transactionMerchant"`
		TransactionPostedDate           time.Time `json:"transactionPostedDate,omitempty"`
		TransactionPostedSequenceNumber int       `json:"transactionPostedSequenceNumber,omitempty"`
		TransactionReferenceID          string    `json:"transactionReferenceId"`
		TransactionState                string    `json:"transactionState"`
		TransactionSubcategoryCode      string    `json:"transactionSubcategoryCode,omitempty"`
		TransactionType                 string    `json:"transactionType,omitempty"`
		TransactionTypeCode             string    `json:"transactionTypeCode,omitempty"`
		ViewOrderURL                    string    `json:"viewOrderUrl,omitempty"`
	} `json:"entries"`
	IsHaMode          bool `json:"isHAMode"`
	IsPartialResponse bool `json:"isPartialResponse"`
	IsVcnRedacted     bool `json:"isVCNRedacted"`
}
