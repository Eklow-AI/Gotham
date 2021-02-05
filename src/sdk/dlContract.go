package sdk

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/Eklow-AI/Gotham/src/models"
)

func getContract(id string) models.ContractProps {
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

	var apiResponse models.RedShirtResp
	json.NewDecoder(resp.Body).Decode(&apiResponse)
	Contract := apiResponse.ContractHistory[0]
	return models.ContractProps{
		ContractAgency:  Contract.ContractAgency,
		FundingAgency:   Contract.FundingAgency,
		Naics:           Contract.Naics,
		Psc:             Contract.Psc,
		SetAside:        Contract.SetAside,
		COSizeSelection: Contract.COSizeSelection,
		CO:              Contract.CO,
		PlaceOfPerf:     Contract.PlaceOfPerf,
		NumOffers:       Contract.NumOffers,
	}
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

func getCContract(cage string, c chan models.ContractProps) {
	c <- getContract(cage)
}
