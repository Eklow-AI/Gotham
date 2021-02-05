package models

import "time"

// RedShirtResp is the struct that binds to the RedShirt API response objects
type RedShirtResp struct {
	NoOfPages            int64          `json:"noofPages"`
	CurrentPage          int64          `json:"currentPage"`
	RecordPerPage        int64          `json:"recordPerPage"`
	AwardDiscovered      *string        `json:"awardDiscovered"`
	TotalValueOfAwards   *string        `json:"totalValueOfAwards"`
	AverageValueOfAwards *string        `json:"averageValueOfAwards"`
	TimeToExecute        float64        `json:"time_to_execute"`
	ContractHistory      []ContractData `json:"listdata"`
}

// ContractData contains actual contract data
type ContractData struct {
	CreatedDate      *time.Time `json:"created_date"`
	SolicitationDate *time.Time `json:"solicitation_date"`
	TotalContractVal float64    `json:"totalcontractvalue,string"`
	NumOffers        int64      `json:"number_of_offers_received,string"`
	VendorName       string     `json:"vendor_name"`
	VendorZip        int64      `json:"phy_zip_code,string"`
	IdvID            string     `json:"idv_id"`
	ContractAgency   string     `json:"contract_agency_name"`
	FundingAgency    string     `json:"funding_agency_name"`
	Naics            string     `json:"naics_code"`
	Psc              string     `json:"product_or_service_code_text"`
	PlaceOfPerf      string     `json:"place_of_performance_zip_code_city"`
	State            string     `json:"state_code_text"`
	COSizeSelection  string     `json:"contracting_officer_business_size_determination_description"`
	CO               string     `json:"last_modified_by"`
	SetAside         string     `json:"type_of_set_aside_description"`
	Cage             string     `json:"cage_code"`
}

// RedShirtQuery is the struct used to create RedShirt API queries
type RedShirtQuery struct {
	Object        string   `json:"object"`
	Version       string   `json:"version"`
	Timeout       int64    `json:"timeout"`
	RecordLimit   int64    `json:"record_limit"`
	Rows          bool     `json:"rows"`
	Totals        bool     `json:"totals"`
	Lists         bool     `json:"lists"`
	SearchFilter  []Filter `json:"searchFilter"`
	RecordPerPage int64    `json:"recordPerPage"`
	CurrentPage   int64    `json:"currentPage"`
	SortFilter    []Filter `json:"sortFilter"`
}

// Filter is the struct used to filter RedShirt API data
type Filter struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
	Order    string `json:"order"`
}
