package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
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
	RecordPerPage int            `json:"recordPerPage"`
	CurrentPage   int            `json:"currentPage"`
	SortFilters   []sortFilter   `json:"sortFilter"`
}

// RSResponse
type RSResponse struct {
	NoOfpages            int          `json:"noofPages"`
	Currentpage          int          `json:"currentPage"`
	Recordperpage        int          `json:"recordPerPage"`
	AwardsDiscovered     int          `json:"awardDiscovered"`
	TotalValueOfAwards   float64      `json:"totalValueOfAwards"`
	AverageValueOfAwards float64      `json:"averageValueOfAwards"`
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

func getRSResponse(query RSRequest) (rsData RSResponse, err error) {
	data, err := json.Marshal(query)
	if err != nil {
		return RSResponse{}, err
	}

	resp, err := client.Post("https://redshirttest.g2xchange.com/wp-json/api/v1/query/?db=MM10TEST", "application/json", bytes.NewBuffer(data))

	if err != nil {
		return RSResponse{}, err
	}
	err = json.NewDecoder(resp.Body).Decode(&rsData)

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		rsData.ListData = []RSContract{}
		return rsData, nil
	}

	if err != nil {
		return RSResponse{}, err
	}
	return rsData, err
}

func getContractFromID(id string) RSContract {
	query := RSRequest{
		Object:      "contracts",
		Version:     "1.2",
		TimeOut:     45000,
		RecordLimit: 10000,
		Rows:        true,
		Totals:      false,
		Lists:       false,
		SearchFilters: []searchFilter{
			{
				Field:    "contract_number",
				Operator: "eq",
				Value:    id,
			},
		},
		RecordPerPage: 1,
		CurrentPage:   1,
		SortFilters: []sortFilter{
			{
				Field: "date_signed",
				Order: "desc",
			},
		},
	}

	data, err := getRSResponse(query)

	if err != nil {
		log.Fatal(err)
	}

	contracts := data.ListData
	if len(contracts) < 1 {
		return RSContract{}
	}
	return contracts[0]
}

func GetContractsFromCage(cage string) []RSContract {
	start := time.Now()
	defer func() {
		fmt.Println("Execution Time: ", time.Since(start))
	}()
	// grab the first contract, create an array the size of all contracts discovered, and then
	// add the first contract to that array
	recordsPerPage := 10
	pages := 6
	contractMap := make([][]RSContract, recordsPerPage*pages)
	wg := sync.WaitGroup{}
	// Get the rest of the contracts concurrently
	for page := 0; page < pages; page++ {
		wg.Add(1)
		go func(currPage int) {
			start := time.Now()
			defer func() {
				fmt.Println("Gorouting Exec: ", time.Since(start))
			}()
			query := RSRequest{
				Object:      "contracts",
				Version:     "1.2",
				TimeOut:     45000,
				RecordLimit: 10000,
				Rows:        true,
				Totals:      true,
				Lists:       false,
				SearchFilters: []searchFilter{
					{
						Field:    "cage_code",
						Operator: "eq",
						Value:    cage,
					},
				},
				RecordPerPage: recordsPerPage,
				CurrentPage:   currPage,
				SortFilters: []sortFilter{
					{
						Field: "date_signed",
						Order: "desc",
					},
				},
			}
			response, err := getRSResponse(query)
			if err != nil {
				log.Fatal(err)
			}
			if len(response.ListData) < 1 {
				fmt.Println("Useless Request")
			}
			contractMap[currPage] = response.ListData
			wg.Done()
		}(page + 1)
	}
	wg.Wait()
	// O(N) time complexity becaused append is O(1)
	contracts := []RSContract{}
	for _, val := range contractMap {
		contracts = append(contracts, val...)
	}
	return contracts
}
