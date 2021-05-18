package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type searchFilter struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type sortFilter struct {
	Field string `json:"field"`
	Order string `json:"order"`
}

// RSRequest is the RedShirt API request object used for queries
type RSRequest struct {
	Object        string         `json:"object"`
	Version       string         `json:"version"`
	TimeOut       int            `json:"timeout"`
	RecordLimit   int            `json:"record_limit"`
	Rows          bool           `json:"rows"`
	Totals        bool           `json:"totals"`
	Lists         bool           `json:"lists"`
	SearchFilters []searchFilter `json:"searchFilter"`
	Recordperpage int            `json:"recordPerPage"`
	Currentpage   int            `json:"currentPage"`
	SortFilters   []sortFilter   `json:"sortFilter"`
}

// RSResponse
type RSResponse struct {
	NoOfpages            int          `json:"noofPages"`
	Currentpage          int          `json:"currentPage"`
	Recordperpage        int          `json:"recordPerPage"`
	Awarddiscovered      int          `json:"awardDiscovered"`
	Totalvalueofawards   int          `json:"totalValueOfAwards"`
	Averagevalueofawards int          `json:"averageValueOfAwards"`
	TimeToExecute        string       `json:"time_to_execute"`
	ListData             []RSContract `json:"listdata"`
}

// RSContract represent the struct of the raw contract data model
// from the RedShirt API
type RSContract struct {
	Cage               string  `json:"cage_code"`
	VendorName         string  `json:"vendor_name"`
	IdvID              string  `json:"idv_id"`
	ID                 string  `json:"contract_piid"`
	ContractAgencyName string  `json:"contract_agency_name"`
	Psc                string  `json:"product_or_service_code_text"`
	Naics              int     `json:"naics_code,string"`
	PhyZipCode         int     `json:"phy_zip_code,string"`
	TotalValue         float64 `json:"totalcontractvalue,string"`
	NumOffers          *int    `json:"number_of_offers_received,string"`
	PlacePerfCity      *string `json:"place_of_performance_zip_code_city"`
	StateCode          *string `json:"state_code_text"`
	SizeSelection      *string `json:"contracting_officer_business_size_determination_description"`
	CO                 *string `json:"last_modified_by"`
	SetAside           *string `json:"type_of_set_aside_description"`
}


func getContractFromID(id string) RSContract {
	query := RSRequest{
		Object:      "contracts",
		Version:     "1.0",
		TimeOut:     45000,
		RecordLimit: 10000,
		Rows:        true,
		Totals:      true,
		Lists:       false,
		SearchFilters: []searchFilter{
			{
				Field:    "contract_number",
				Operator: "eq",
				Value:    id,
			},
		},
		Recordperpage: 1,
		Currentpage:   1,
		SortFilters: []sortFilter{
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
	resp, err := client.Post(fmt.Sprintf("%s%s", os.Getenv("RS_URI"), "/wp-json/api/v1/query/?db=MM10TEST"), "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	var rsData RSResponse
	json.NewDecoder(resp.Body).Decode(&rsData)
	contracts := rsData.ListData
	if len(contracts) < 1 {
		return RSContract{}
	}
	return contracts[0]
}
