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
	emptystr := ""
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
				SetAside:           "NO SET ASIDE USED.",
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
				SetAside:           &emptystr,
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
				SetAside:           "NO SET ASIDE USED.",
			},
		},
	}
	for num, test := range passToContractTests {
		if output := PassToProfile(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output not equal to expected struct for test %d", num)
		}
	}
}

type calcPercentBreakdownTest struct {
	arg1     []string
	expected map[string]float64
}

var calcPercentBreakdownTests = []calcPercentBreakdownTest{
	{[]string{"agency1", "agency2", "agency3", "agency4"},map[string]float64{"agency1": 0.25, "agency2": 0.25, "agency3": 0.25, "agency4": 0.25},},
	{[]string{"DoD"},map[string]float64{"DoD": 1.0},},
	{[]string{"DoD","VA"},map[string]float64{"DoD": 0.5, "VA": 0.5},},
	{[]string{},map[string]float64{},},
}

func TestPercentBreakdown(t *testing.T) {
	for num, test := range calcPercentBreakdownTests {
		if output := CalcPercentBreakdown(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output not equal to expected map for test %d", num)
		} 
	}
}
