package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

type contractTest struct {
	jsonBlob []byte
	isNull   bool
}

var contractTests = []contractTest{
	{
		[]byte(`{
			"created_date": "2020-05-14 00:00:00",
			"map_fpds_id": "1609216",
			"vendor_name": "VISUAL CONNECTIONS L.L.C.",
			"phy_zip_code": "21244",
			"idv_id": "VA26315A0050",
			"contract_piid": "VA26316J0726",
			"contract_mod_number": "0",
			"signed_date": "2016-07-20 00:00:00",
			"totalcontractvalue": "250000.00",
			"totalspendtodate": "1983.46",
			"total_base_and_all_options_value": "250000.00",
			"total_action_obligation": "1983.46",
			"number_of_offers_received": "1",
			"contract_agency_name": "VETERANS AFFAIRS, DEPARTMENT OF",
			"contract_agency_id": "3600",
			"funding_agency_name": "VETERANS AFFAIRS, DEPARTMENT OF",
			"funding_agency_id": "3600",
			"funding_office_name": "636-NEBRASKA WESTERN-IOWA (00636)",
			"funding_office_id": "36C636",
			"contractingagency": "VETERANS AFFAIRS, DEPARTMENT OF (3600)",
			"fundingagency": "VETERANS AFFAIRS, DEPARTMENT OF (3600)",
			"fundingoffice": "636-NEBRASKA WESTERN-IOWA (00636) (36C636)",
			"ultimate_completion_date": "2020-05-14 00:00:00",
			"naics_code": "541511",
			"idvtype": "BPA CALL",
			"principal_naics_code_description": "CUSTOM COMPUTER PROGRAMMING SERVICES",
			"product_or_service_code_text": "J070",
			"product_or_service_code_description": "MAINT/REPAIR/REBUILD OF EQUIPMENT- ADP EQUIPMENT/SOFTWARE/SUPPLIES/SUPPORT EQUIPMENT",
			"place_of_performance_zip_code_city": "DES MOINES",
			"state_code_text": "IA",
			"contracting_officer_business_size_determination_description": "SMALL BUSINESS",
			"last_modified_by": "IDV_CORRECT",
			"type_of_set_aside_description": "SERVICE DISABLED VETERAN OWNED SMALL BUSINESS SET-ASIDE",
			"extent_competed_description": "FULL AND OPEN COMPETITION",
			"multiple_or_single_award_idc_description": "SINGLE AWARD",
			"description_of_contract_requirement": "IGF::OT::IGF\nCABLING SERVICE FOR DES MOINES, IA VETERANS ADMINISTRATION AND ASSIGNED COMMUNITY BASED OUTPATIENT CLINIC LOCATONS.",
			"dtc": "-366",
			"psc_code_string": "AB12",
			"cage_code": "6ZP36",
			"duns": "808543123",
			"reason_for_modification": "M: OTHER ADMINISTRATIVE ACTION",
			"solicitation_id": "RFPCMS2016SPARC",
			"solicitation_procedures_description": "SUBJECT TO MULTIPLE AWARD FAIR OPPORTUNITY",
			"mods_array": false
		}`),
		false,
	},
	{
		[]byte(`{
			"created_date": "2020-05-14 00:00:00",
			"map_fpds_id": "1609216",
			"vendor_name": "VISUAL CONNECTIONS L.L.C.",
			"phy_zip_code": "21244",
			"idv_id": "VA26315A0050",
			"contract_piid": "VA26316J0726",
			"contract_mod_number": "0",
			"signed_date": "2016-07-20 00:00:00",
			"totalcontractvalue": "250000.00",
			"totalspendtodate": "1983.46",
			"total_base_and_all_options_value": "250000.00",
			"total_action_obligation": "1983.46",
			"number_of_offers_received": null,
			"contract_agency_name": "VETERANS AFFAIRS, DEPARTMENT OF",
			"contract_agency_id": "3600",
			"funding_agency_name": "VETERANS AFFAIRS, DEPARTMENT OF",
			"funding_agency_id": "3600",
			"funding_office_name": "636-NEBRASKA WESTERN-IOWA (00636)",
			"funding_office_id": "36C636",
			"contractingagency": "VETERANS AFFAIRS, DEPARTMENT OF (3600)",
			"fundingagency": "VETERANS AFFAIRS, DEPARTMENT OF (3600)",
			"fundingoffice": "636-NEBRASKA WESTERN-IOWA (00636) (36C636)",
			"ultimate_completion_date": "2020-05-14 00:00:00",
			"naics_code": "541511",
			"idvtype": "BPA CALL",
			"principal_naics_code_description": "CUSTOM COMPUTER PROGRAMMING SERVICES",
			"product_or_service_code_text": "J070",
			"product_or_service_code_description": "MAINT/REPAIR/REBUILD OF EQUIPMENT- ADP EQUIPMENT/SOFTWARE/SUPPLIES/SUPPORT EQUIPMENT",
			"place_of_performance_zip_code_city": null,
			"state_code_text": null,
			"contracting_officer_business_size_determination_description": null,
			"last_modified_by": null,
			"type_of_set_aside_description": null,
			"extent_competed_description": "FULL AND OPEN COMPETITION",
			"multiple_or_single_award_idc_description": "SINGLE AWARD",
			"description_of_contract_requirement": "IGF::OT::IGF\nCABLING SERVICE FOR DES MOINES, IA VETERANS ADMINISTRATION AND ASSIGNED COMMUNITY BASED OUTPATIENT CLINIC LOCATONS.",
			"dtc": "-366",
			"psc_code_string": "AB12",
			"cage_code": "6ZP36",
			"duns": "808543123",
			"reason_for_modification": "M: OTHER ADMINISTRATIVE ACTION",
			"solicitation_id": "RFPCMS2016SPARC",
			"solicitation_procedures_description": "SUBJECT TO MULTIPLE AWARD FAIR OPPORTUNITY",
			"mods_array": false
		}`),
		true,
	},
}

func TestRSContract(t *testing.T) {

	for num, test := range contractTests {
		contract := RSContract{}
		err := json.Unmarshal(test.jsonBlob, &contract)
		// throw an err if there problems unmarshaling
		if err != nil {
			t.Errorf("%q", err)
			return
		}
		// test fields which are never expected to be null
		if reflect.TypeOf(contract.ContractAgencyName).Kind() != reflect.String {
			t.Errorf("Expected: %q Got: %q  Test: %q", reflect.String, reflect.TypeOf(contract.ContractAgencyName).Kind(), num)
		}
		if reflect.TypeOf(contract.VendorName).Kind() != reflect.String {
			t.Errorf("Expected: %q Got: %q  Test: %q", reflect.String, reflect.TypeOf(contract.VendorName).Kind(), num)
		}
		if reflect.TypeOf(contract.Cage).Kind() != reflect.String {
			t.Errorf("Expected: %q Got: %q  Test: %q", reflect.String, reflect.TypeOf(contract.Cage).Kind(), num)
		}
		if reflect.TypeOf(contract.PhyZipCode).Kind() != reflect.Int {
			t.Errorf("Expected: %q Got: %q  Test: %d", reflect.Int, reflect.TypeOf(contract.PhyZipCode).Kind(), num)
		}
		if reflect.TypeOf(contract.IdvID).Kind() != reflect.String {
			t.Errorf("Expected: %q Got: %q  Test: %d", reflect.String, reflect.TypeOf(contract.IdvID).Kind(), num)
		}
		if reflect.TypeOf(contract.ID).Kind() != reflect.String {
			t.Errorf("Expected: %q Got: %q  Test: %d", reflect.String, reflect.TypeOf(contract.ID).Kind(), num)
		}
		if reflect.TypeOf(contract.Psc).Kind() != reflect.String {
			t.Errorf("Expected: %q Got: %q  Test: %d", reflect.String, reflect.TypeOf(contract.Psc).Kind(), num)
		}
		if reflect.TypeOf(contract.Naics).Kind() != reflect.Int {
			t.Errorf("Expected: %q Got: %q  Test: %d", reflect.Int, reflect.TypeOf(contract.Psc).Kind(), num)
		}
		if reflect.TypeOf(contract.TotalValue).Kind() != reflect.Float64 {
			t.Errorf("Expected: %q Got: %q  Test: %d", reflect.Float64, reflect.TypeOf(contract.TotalValue).Kind(), num)
		}
		// Test possible null fields
		if !test.isNull && reflect.TypeOf(*contract.SetAside).Kind() != reflect.String {
			t.Errorf("Expected: %q Got: %q  Test: %q", reflect.String, reflect.TypeOf(*contract.SetAside).Kind(), num)
		}
		if !test.isNull && reflect.TypeOf(*contract.NumOffers).Kind() != reflect.Int {
			t.Errorf("Expected: %q Got: %q  Test: %d", reflect.Int, reflect.TypeOf(*contract.NumOffers).Kind(), num)
		}
		if !test.isNull && reflect.TypeOf(*contract.PlacePerfCity).Kind() != reflect.String {
			t.Errorf("Expected: %q Got: %q  Test: %d", reflect.String, reflect.TypeOf(*contract.PlacePerfCity).Kind(), num)
		}
		if !test.isNull && reflect.TypeOf(*contract.SizeSelection).Kind() != reflect.String {
			t.Errorf("Expected: %q Got: %q  Test: %d", reflect.String, reflect.TypeOf(*contract.SizeSelection).Kind(), num)
		}
		if !test.isNull && reflect.TypeOf(*contract.StateCode).Kind() != reflect.String {
			t.Errorf("Expected: %q Got: %q  Test: %d", reflect.String, reflect.TypeOf(*contract.StateCode).Kind(), num)
		}
		if !test.isNull && reflect.TypeOf(*contract.CO).Kind() != reflect.String {
			t.Errorf("Expected: %q Got: %q  Test: %d", reflect.String, reflect.TypeOf(*contract.CO).Kind(), num)
		}
		// test when it is a null test
		if test.isNull {
			if contract.SetAside != nil {
				t.Errorf("Expected: nil Got: %q  Test: %q", reflect.TypeOf(*contract.SetAside).Kind(), num)
			}
			if contract.NumOffers != nil {
				t.Errorf("Expected: nil Got: %q  Test: %d", reflect.TypeOf(*contract.NumOffers).Kind(), num)
			}
			if contract.PlacePerfCity != nil {
				t.Errorf("Expected: nil Got: %q  Test: %d", reflect.TypeOf(*contract.PlacePerfCity).Kind(), num)
			}
			if contract.SizeSelection != nil {
				t.Errorf("Expected: nil Got: %q  Test: %d", reflect.TypeOf(*contract.SizeSelection).Kind(), num)
			}
			if contract.StateCode != nil {
				t.Errorf("Expected: nil Got: %q  Test: %d", reflect.TypeOf(*contract.StateCode).Kind(), num)
			}
			if contract.CO != nil {
				t.Errorf("Expected: nil Got: %q  Test: %d", reflect.TypeOf(*contract.CO).Kind(), num)
			}
		}
	}
}
