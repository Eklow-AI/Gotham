package models

// RedShirtResp is the struct that binds to the RedShirt API response objects
type RedShirtResp struct {
	Status               status                   `json:"status"`
	NoOfPages            int64                    `json:"noofPages"`
	CurrentPage          int64                    `json:"currentPage"`
	RecordPerPage        int64                    `json:"recordPerPage"`
	AwardDiscovered      string                   `json:"awardDiscovered"`
	TotalValueOfAwards   string                   `json:"totalValueOfAwards"`
	AverageValueOfAwards string                   `json:"averageValueOfAwards"`
	TimeToExecute        float64                  `json:"time_to_execute"`
	ListData             []map[string]interface{} `json:"listdata"`
}

// RedShirtQuery is the struct used to create RedShirt API queries
type RedShirtQuery struct {
	Object        string   `json:"object"`
	Version       string   `json:"versions"`
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

type status struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type info struct {
	Database      string `json:"database"`
	Software      string `json:"software"`
	SAMUpdate     string `json:"SAM Update"`
	FDPSUpdate    string `json:"FPDS Update"`
	TotalRecords  int64  `json:"totalRecords2"`
	CacheID       string `json:"cache_id"`
	CacheIDLength int64  `json:"cache_id_length"`
	UsingCache    bool   `json:"Using Cache"`
}
