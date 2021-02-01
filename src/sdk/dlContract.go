package sdk

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/Eklow-AI/Gotham/src/models"
)

func getContract(id string) (contract models.ContractProps) {
	query := models.RedShirtQuery{
		Object:      "contracts",
		Version:     "1.0",
		Timeout:     45000,
		RecordLimit: 10000,
		Rows:        true,
		Totals:      true,
		Lists:       false,
		SearchFilter: []models.Filter{
			{
				Field:    "contract_name", //placeholder until I find out how to look up by ID
				Operator: "tsquery",
				Value:    id,
			},
		},
		RecordPerPage: 1,
		CurrentPage:   1,
		SortFilter: []models.Filter{
			{
				Field: "date_signed",
				Order: "desc",
			},
		},
	}
	data, err := json.Marshal(query)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Post("https://redshirttest.g2xchange.com/wp-json/api/v1/query/?db=MM10TEST",
		"application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	var dataResp models.RedShirtResp
	json.NewDecoder(resp.Body).Decode(&dataResp)
	// TODO: what if any value comes in nil? This could break with type assertions
	contract = models.ContractProps{
		ContractAgency:  dataResp.ListData[0]["contract_agency_name"].(string),
		FundingAgency:   dataResp.ListData[0]["funding_agency_name"].(string),
		Naics:           dataResp.ListData[0]["naics_code"].(string),
		Psc:             dataResp.ListData[0]["product_or_service_code_text"].(string),
		SetAside:        dataResp.ListData[0]["type_of_set_aside_description"].(string),
		COSizeSelection: dataResp.ListData[0]["contracting_officer_business_size_determination_description"].(string),
		CO:              dataResp.ListData[0]["last_modified_by"].(string),
		PlaceOfPerf:     dataResp.ListData[0]["place_of_performance_zip_code_city"].(string),
		NumOffers:       dataResp.ListData[0]["number_of_offers_received"].(int64),
	}
	return contract
}

func getContractProfile(contract models.ContractProps, vendor models.VendorProfile) models.ContractProfile {
	return models.ContractProfile{
		ContractAgency:  vendor.ContractAgency[contract.ContractAgency],
		FundingAgency:   vendor.FundingAgency[contract.FundingAgency],
		Naics:           vendor.Naics[contract.Naics],
		Psc:             vendor.Psc[contract.Psc],
		SetAside:        vendor.SetAsides[contract.SetAside],
		COSizeSelection: vendor.COSizeSelection[contract.COSizeSelection],
		CO:              vendor.COs[contract.CO],
		PlaceOfPerf:     vendor.PlacesOfPerf[contract.PlaceOfPerf],
		NumOffers:       contract.NumOffers,
	}
}
