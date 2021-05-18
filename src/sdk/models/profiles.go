package models

// ContractProfile is the standard struct used by the SDK
// to analyze compatability scores
type ContractProfile struct {
	PhyZipCode         int
	Naics              int
	NumOffers          int
	TotalValue         float64
	Cage               string
	VendorName         string
	IdvID              string
	ID                 string
	Psc                string
	ContractAgencyName string
	PlacePerfCity      string
	StateCode          string
	SizeSelection      string
	CO                 string
	SetAside           string
}

// VendorProfile is the standard struct used by the SDK
// to represent vendor profiles
type VendorProfile struct {
	Zip             int
	Cage            string
	Name            string
	AvgContractSize float64
	Naics           map[int64]float64
	Psc             map[string]float64
	COs             map[string]float64
	SetAsides       map[string]float64
}

// PassToCProfile accepts an RDContract object and translates it to
// a ContractProfile.
func PassToCProfile(contract RSContract) ContractProfile {
	profile := ContractProfile{}
	// Populate all standard fields to the rest of the profile
	profile.Cage = contract.Cage
	profile.VendorName = contract.VendorName
	profile.IdvID = contract.IdvID
	profile.ID = contract.ID
	profile.ContractAgencyName = contract.ContractAgencyName
	profile.Psc = contract.Psc
	profile.Naics = contract.Naics
	profile.PhyZipCode = contract.PhyZipCode
	profile.TotalValue = contract.TotalValue
	// Check possible nil values and replace them
	profile.NumOffers = 1
	if contract.NumOffers != nil {
		profile.NumOffers = *contract.NumOffers
	}
	profile.PlacePerfCity = "unknown"
	if contract.PlacePerfCity != nil {
		profile.PlacePerfCity = *contract.PlacePerfCity
	}
	profile.StateCode = "unknown"
	if contract.StateCode != nil {
		profile.StateCode = *contract.StateCode
	}
	profile.SizeSelection = "unknown"
	if contract.SizeSelection != nil {
		profile.SizeSelection = *contract.SizeSelection
	}
	profile.CO = "unknown"
	if contract.CO != nil {
		profile.CO = *contract.CO
	}
	profile.SetAside = "NO SET ASIDE USED."
	if contract.SetAside != nil {
		profile.SetAside = *contract.SetAside
		if *contract.SetAside == "" {
			profile.SetAside = "NO SET ASIDE USED."
		}
	}
	return profile
}

func PassToVProfile(contracts []RSContract) VendorProfile {
	vendor := VendorProfile{}

}

func CalcPercentBreakdown(items []string) map[string]float64 {
	breakdown := map[string]float64{}
	// First, populate the map with the count of each item
	for _, item := range items {
		if _, ok := breakdown[item]; ok {
			breakdown[item] += 1.0
			continue
		}
		breakdown[item] = 1
	}
	// Then calculate the percentages of each
	totalItems := float64(len(items))
	for key, val := range breakdown {
		breakdown[key] = val / totalItems
	}
	return breakdown
}
