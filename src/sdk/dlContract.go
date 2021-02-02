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
		Version:     "1.2",
		Timeout:     45000,
		RecordLimit: 10000,
		Rows:        true,
		Totals:      true,
		Lists:       false,
		SearchFilter: []models.Filter{
			{
				Field:    "contract_number",
				Operator: "eq",
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
	if err != nil {
		log.Fatal(err)
	}

	contract = models.ContractProps{
		ContractAgency:  dataResp.ContractData[0].ContractAgency,
		FundingAgency:   dataResp.ContractData[0].FundingAgency,
		Naics:           dataResp.ContractData[0].Naics,
		Psc:             dataResp.ContractData[0]["product_or_service_code_text"],
		SetAside:        dataResp.ContractData[0]["type_of_set_aside_description"],
		COSizeSelection: dataResp.ContractData[0]["contracting_officer_business_size_determination_description"],
		CO:              dataResp.ContractData[0]["last_modified_by"],
		PlaceOfPerf:     dataResp.ContractData[0]["place_of_performance_zip_code_city"],
		NumOffers:       numOffers,
	}
	return contract
}

func recoverAssertion(input interface{}) string {
	output, ok := input.(string)
	if !ok {
		return "unknown"
	}
	return output
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
