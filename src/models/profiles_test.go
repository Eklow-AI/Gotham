package models

import (
	"reflect"
	"testing"
)

type passToContractTest struct {
	arg1     RSContract
	expected ContractProfile
}

func TestPassToContractProfile(t *testing.T) {
	// Test pointer values
	offers := 3
	zip := "Miami"
	state := "FL"
	ss := "small biz"
	co := "Bob Fries"
	setaside := "diversity owned"
	var passToContractTests = []passToContractTest{
		{
			RSContract{
				Cage:               "AAAA",
				VendorName:         "Eklow AI",
				PhyZipCode:         33328,
				IdvID:              "someidstuf13412",
				ID:                 "another123ID",
				Psc:                "PD13",
				Naics:              411245,
				ContractAgencyName: "Some Agency over there",
				TotalValue:         1000000.00,
				NumOffers:          &offers,
				PlacePerfCity:      &zip,
				StateCode:          &state,
				SizeSelection:      &ss,
				CO:                 &co,
				SetAside:           &setaside,
			},
			ContractProfile{
				Cage:               "AAAA",
				VendorName:         "Eklow AI",
				PhyZipCode:         33328,
				IdvID:              "someidstuf13412",
				ID:                 "another123ID",
				Psc:                "PD13",
				Naics:              411245,
				ContractAgencyName: "Some Agency over there",
				TotalValue:         1000000.00,
				NumOffers:          offers,
				PlacePerfCity:      zip,
				StateCode:          state,
				SizeSelection:      ss,
				CO:                 co,
				SetAside:           setaside,
			},
		},
		{
			RSContract{
				Cage:               "AAAA",
				VendorName:         "Eklow AI",
				PhyZipCode:         33328,
				IdvID:              "someidstuf13412",
				ID:                 "another123ID",
				Psc:                "PD13",
				Naics:              411245,
				ContractAgencyName: "Some Agency over there",
				TotalValue:         1000000.00,
				NumOffers:          nil,
				PlacePerfCity:      nil,
				StateCode:          nil,
				SizeSelection:      nil,
				CO:                 nil,
				SetAside:           nil,
			},
			ContractProfile{
				Cage:               "AAAA",
				VendorName:         "Eklow AI",
				PhyZipCode:         33328,
				IdvID:              "someidstuf13412",
				ID:                 "another123ID",
				Psc:                "PD13",
				Naics:              411245,
				ContractAgencyName: "Some Agency over there",
				TotalValue:         1000000.00,
				NumOffers:          1,
				PlacePerfCity:      "unknown",
				StateCode:          "unknown",
				SizeSelection:      "unknown",
				CO:                 "unknown",
				SetAside:           "unknown",
			},
		},
	}
	for num, test := range passToContractTests {
		if output := PassToProfile(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output not equal to expected struct for test %d", num)
		}
	}
}
