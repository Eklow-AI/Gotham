package models

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