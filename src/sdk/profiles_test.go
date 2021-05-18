package sdk

import (
	"reflect"
	"testing"
)

type passToContractTest struct {
	arg1     RSContract
	expected ContractProfile
}

func TestPassContractToProfile(t *testing.T) {
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
		if output := PassToCProfile(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output not equal to expected struct for test %d", num)
		}
	}
}

type calcPercentBreakdownTest struct {
	arg1     []string
	expected map[string]float64
}

var calcPercentBreakdownTests = []calcPercentBreakdownTest{
	{[]string{"agency1", "agency2", "agency3", "agency4"}, map[string]float64{"agency1": 0.25, "agency2": 0.25, "agency3": 0.25, "agency4": 0.25}},
	{[]string{"DoD"}, map[string]float64{"DoD": 1.0}},
	{[]string{"DoD", "VA"}, map[string]float64{"DoD": 0.5, "VA": 0.5}},
	{[]string{}, map[string]float64{}},
}

func TestPercentBreakdown(t *testing.T) {
	for num, test := range calcPercentBreakdownTests {
		if output := calcPercentBreakdown(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output not equal to expected map for test %d", num)
		}
	}
}

type passToVProfileTest struct {
	arg1     []ContractProfile
	expected VendorProfile
}

var passToVProfileTests = []passToVProfileTest{
	{
		[]ContractProfile{},
		VendorProfile{},
	},
	{
		[]ContractProfile{
			{
				PhyZipCode:         55555,
				Naics:              1234567,
				NumOffers:          1,
				TotalValue:         100.0,
				Cage:               "ABCD",
				VendorName:         "Eklow AI",
				IdvID:              "someid1234",
				ID:                 "anotherid123$",
				Psc:                "D233",
				ContractAgencyName: "DoD",
				PlacePerfCity:      "Caracas",
				StateCode:          "VZLA",
				SizeSelection:      "Small Biz Size",
				CO:                 "mckerzier@hotmail.com",
				SetAside:           "minority owned",
			},
		},
		VendorProfile{
			Name:            "Eklow AI",
			Zip:             55555,
			Cage:            "ABCD",
			AvgContractSize: 100.0,
			Naics:           map[int]float64{1234567: 1.0},
			Psc:             map[string]float64{"D233": 1.0},
			COs:             map[string]float64{"mckerzier@hotmail.com": 1.0},
			SetAsides:       map[string]float64{"minority owned": 1.0},
		},
	},
	{
		[]ContractProfile{
			{
				PhyZipCode:         55555,
				Naics:              1234567,
				NumOffers:          1,
				TotalValue:         200.0,
				Cage:               "ABCD",
				VendorName:         "Eklow AI",
				IdvID:              "someid1234",
				ID:                 "anotherid123$",
				Psc:                "D233",
				ContractAgencyName: "DoD",
				PlacePerfCity:      "Caracas",
				StateCode:          "VZLA",
				SizeSelection:      "Small Biz Size",
				CO:                 "mckerzier@hotmail.com",
				SetAside:           "minority owned",
			},
			{
				PhyZipCode:         55555,
				Naics:              1234567,
				NumOffers:          1,
				TotalValue:         300.0,
				Cage:               "ABCD",
				VendorName:         "Eklow AI",
				IdvID:              "someid1234",
				ID:                 "anotherid123$",
				Psc:                "L222",
				ContractAgencyName: "DoD",
				PlacePerfCity:      "Caracas",
				StateCode:          "VZLA",
				SizeSelection:      "Small Biz Size",
				CO:                 "mckerzier@hotmail.com",
				SetAside:           "minority owned",
			},
			{
				PhyZipCode:         55555,
				Naics:              -100,
				NumOffers:          1,
				TotalValue:         400.0,
				Cage:               "ABCD",
				VendorName:         "Eklow AI",
				IdvID:              "someid1234",
				ID:                 "anotherid123$",
				Psc:                "A111",
				ContractAgencyName: "DoD",
				PlacePerfCity:      "Caracas",
				StateCode:          "VZLA",
				SizeSelection:      "Small Biz Size",
				CO:                 "atlas@eklow.com",
				SetAside:           "NO SET ASIDE USED.",
			},
			{
				PhyZipCode:         55555,
				Naics:              999999999999,
				NumOffers:          1,
				TotalValue:         500.0,
				Cage:               "ABCD",
				VendorName:         "Eklow AI",
				IdvID:              "someid1234",
				ID:                 "anotherid123$",
				Psc:                "Z000",
				ContractAgencyName: "DoD",
				PlacePerfCity:      "Caracas",
				StateCode:          "VZLA",
				SizeSelection:      "Small Biz Size",
				CO:                 "atlas@eklow.com",
				SetAside:           "NO SET ASIDE USED.",
			},
		},
		VendorProfile{
			Name:            "Eklow AI",
			Zip:             55555,
			Cage:            "ABCD",
			AvgContractSize: 350.0,
			Naics:           map[int]float64{1234567: 0.5, -100: 0.25, 999999999999: 0.25},
			Psc:             map[string]float64{"D233": 0.25, "A111": 0.25, "L222": 0.25, "Z000": 0.25},
			COs:             map[string]float64{"mckerzier@hotmail.com": 0.5, "atlas@eklow.com": 0.5},
			SetAsides:       map[string]float64{"minority owned": 0.5, "NO SET ASIDE USED.": 0.5},
		},
	},
}

func TestPassToVProfile(t *testing.T) {
	for num, test := range passToVProfileTests {
		output := PassToVProfile(test.arg1)
		if output.Name != test.expected.Name {
			t.Errorf("Vendor name does not match expected %d", num)
		}
		if output.Cage != test.expected.Cage {
			t.Errorf("Vendor Cage code does not match expected %d", num)
		}
		if output.Zip != test.expected.Zip {
			t.Errorf("Vendor Cage code does not match expected %d", num)
		}
		if output.AvgContractSize != test.expected.AvgContractSize {
			t.Errorf("Output avg contract size does not match expected %d", num)
		}
		if !reflect.DeepEqual(output.Naics, test.expected.Naics) {
			t.Errorf("Vendor Naics slice does not match expected %d", num)
		}
		if !reflect.DeepEqual(output.Psc, test.expected.Psc) {
			t.Errorf("Vendor Pscs slice does not match expected %d", num)
		}
		if !reflect.DeepEqual(output.COs, test.expected.COs) {
			t.Errorf("Vendor COs slice does not match expected %d", num)
		}
		if !reflect.DeepEqual(output.SetAsides, test.expected.SetAsides) {
			t.Errorf("Vendor set asides slice does not match expected %d", num)
		}
	}
}
