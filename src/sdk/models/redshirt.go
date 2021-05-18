package models

type searchFilter struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    int    `json:"value"`
}

type sortFilter struct {
	Field string `json:"field"`
	Order string `json:"order"`
}

// RSRequest is the RedShirt API request object used for queries
type RSRequest struct {
	Object        string         `json:"object"`
	Version       string         `json:"version"`
	Timeout       int            `json:"timeout"`
	RecordLimit   int            `json:"record_limit"`
	Rows          bool           `json:"rows"`
	Totals        bool           `json:"totals"`
	Lists         bool           `json:"lists"`
	Searchfilter  []searchFilter `json:"searchFilter"`
	Recordperpage int            `json:"recordPerPage"`
	Currentpage   int            `json:"currentPage"`
	Sortfilter    []sortFilter   `json:"sortFilter"`
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
	Listdata             []RSContract `json:"listdata"`
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
