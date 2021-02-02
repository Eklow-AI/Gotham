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
	ContractData         []ContractData `json:"listdata"`
}

// ContractData contains actual contract data
type ContractData struct {
	CreatedDate                         *time.Time `json:"created_date"`
	MapFpdsID                           *string    `json:"map_fpds_id"`
	VendorName                          *string    `json:"vendor_name"`
	PhyZipCode                          *string    `json:"phy_zip_code"`
	IdvID                               *string    `json:"idv_id"`
	ContractPiid                        *string    `json:"contract_piid"`
	ContractModNumber                   *string    `json:"contract_mod_number"`
	SignedDate                          *time.Time `json:"signed_date"`
	TotalContractVal                    *float64   `json:"totalcontractvalue,string"`
	TotalSpendToDate                    *float64   `json:"totalspendtodate,string"`
	TotalBaseAndAllOptionsValue         *string    `json:"total_base_and_all_options_value"`
	TotalActionObligation               *string    `json:"total_action_obligation"`
	NumOffers                           *int64     `json:"number_of_offers_received,string"`
	ContractAgency                      *string    `json:"contract_agency_name"`
	ContractAgencyID                    *string    `json:"contract_agency_id"`
	FundingAgency                       *string    `json:"funding_agency_name"`
	FundingAgencyID                     *string    `json:"funding_agency_id"`
	FundingOffice                       *string    `json:"funding_office_name"`
	FundingOfficeID                     *string    `json:"funding_office_id"`
	UltimateCompletionDate              *string    `json:"ultimate_completion_date"`
	Naics                               *string    `json:"naics_code"`
	IdvType                             *string    `json:"idvtype"`
	PrincipalNaicsCodeDescription       *string    `json:"principal_naics_code_description"`
	Psc                                 *string    `json:"product_or_service_code_text"`
	ProductOrServiceCodeDescription     *string    `json:"product_or_service_code_description"`
	PlaceOfPerformanceZipCodeCity       *string    `json:"place_of_performance_zip_code_city"`
	State                               *string    `json:"state_code_text"`
	COBizSize                           *string    `json:"contracting_officer_business_size_determination_description"`
	CO                                  *string    `json:"last_modified_by"`
	TypeOfSetAsideDescription           *string    `json:"type_of_set_aside_description"`
	ExtentCompetedDescription           *string    `json:"extent_competed_description"`
	MultipleOrSingleAwardIdcDescription *string    `json:"multiple_or_single_award_idc_description"`
	DescriptionOfContractRequirement    *string    `json:"description_of_contract_requirement"`
	Dtc                                 *string    `json:"dtc"`
	PscCodeString                       *string    `json:"psc_code_string"`
	Cage                                *string    `json:"cage_code"`
	Duns                                *string    `json:"duns"`
	ReasonForModification               *string    `json:"reason_for_modification"`
	SolicitationID                      *string    `json:"solicitation_id"`
	SolicitationProceduresDescription   *string    `json:"solicitation_procedures_description"`
	SolicitationDate                    *time.Time `json:"solicitation_date"`
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
