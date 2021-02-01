package models

// VendorProfile describes the % frequency of contracts across several params
type VendorProfile struct {
	Name            string
	Cage            string
	ContractAgency  map[string]float64
	FundingAgency   map[string]float64
	Naics           map[string]float64
	Psc             map[string]float64
	SetAsides       map[string]float64
	COSizeSelection map[string]float64
	COs             map[string]float64
	PlacesOfPerf    map[string]float64
	Zip             int64
}

// ContractProfile describes an opportunity for scoring
type ContractProfile struct {
	ContractAgency  float64 // d
	FundingAgency   float64
	Naics           float64 // d
	Psc             float64 // d
	SetAside        float64
	COSizeSelection float64
	CO              float64 // d
	PlaceOfPerf     float64
	NumOffers       int64 // d
}

// ContractProps describes the properties of a contract
type ContractProps struct {
	ContractAgency  string
	FundingAgency   string
	Naics           string
	Psc             string
	SetAside        string
	COSizeSelection string
	CO              string
	PlaceOfPerf     string
	NumOffers       int64
}
