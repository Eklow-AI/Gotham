package sdk

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
	Naics           map[int]float64
	Psc             map[string]float64
	COs             map[string]float64
	SetAsides       map[string]float64
}

// PassToCProfile accepts an RDContract object and translates it to
// a ContractProfile.
func PassRSToCProfile(contract RSContract) ContractProfile {
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

func PassRSToVProfile(contracts []ContractProfile) VendorProfile {
	vendor := VendorProfile{}
	if len(contracts) < 1 {
		return vendor
	}
	// Get data constants of the vendor from their first contract
	vendor.Cage = contracts[0].Cage
	vendor.Name = contracts[0].VendorName
	vendor.Zip = contracts[0].PhyZipCode
	// Calculate average contract size
	totalValue := 0.0
	for _, contract := range contracts {
		totalValue += contract.TotalValue
	}
	vendor.AvgContractSize = totalValue / float64(len(contracts))
	// Generate slices representing all the Psc, Naics, COs, and set-asides
	// that have been used by the vendor
	Pscs := []string{}
	for _, contract := range contracts {
		Pscs = append(Pscs, contract.Psc)
	}
	Naics := []int{}
	for _, contract := range contracts {
		Naics = append(Naics, contract.Naics)
	}
	COs := []string{}
	for _, contract := range contracts {
		COs = append(COs, contract.CO)
	}
	ss := []string{}
	for _, contract := range contracts {
		ss = append(ss, contract.SetAside)
	}
	// Calculate the percentage breakdown from it
	vendor.COs = calcPercentBreakdown(COs)
	vendor.SetAsides = calcPercentBreakdown(ss)
	vendor.Psc = calcPercentBreakdown(Pscs)
	// Handle calculate the percentage breakdown for the
	// Naics slice since its int and not string based
	breakdown := map[int]float64{}
	for _, item := range Naics {
		if _, ok := breakdown[item]; ok {
			breakdown[item] += 1.0
			continue
		}
		breakdown[item] = 1
	}
	totalItems := float64(len(Naics))
	for key, val := range breakdown {
		breakdown[key] = val / totalItems
	}
	vendor.Naics = breakdown
	return vendor
}

func calcPercentBreakdown(items []string) map[string]float64 {
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

func getCProfileFromRS(id string) ContractProfile {
	rsContract := getContractFromID(id)
	contract := PassRSToCProfile(rsContract)
	return contract
}
