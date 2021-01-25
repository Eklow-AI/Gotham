package models

// VendorProfile describes the % frequency of contracts across several params
type VendorProfile struct {
	Name            string
	Cage            string
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
	FundingAgency   float64
	Naics           float64
	Psc             float64
	SetAside        float64
	COSizeSelection float64
	CO              float64
	PlaceOfPerf     float64
	NumOffers       int64
}

// ContractProps describes the properties of a contract
type ContractProps struct {
	FundingAgency   string
	Naics           string
	Psc             string
	SetAside        string
	COSizeSelection string
	CO              string
	PlaceOfPerf     string
	NumOffers       int64
}
