package models

type ContractProfile struct {
	Cage               string
	VendorName         string
	PhyZipCode         int
	IdvID              string
	ID                 string
	Psc                string
	ContractAgencyName string
	Naics              int
	TotalValue         float64
	NumOffers          int
	PlacePerfCity      string
	StateCode          string
	SizeSelection      string
	CO                 string
	SetAside           string
}

type VendorProfile struct {
	Cage      string
	Name      string
	Zip       int
	Psc       map[int64]float64
	Naics     map[int64]float64
	COs       map[string]float64
	SetAsides map[string]float64
}

func PassToProfile(contract RSContract) ContractProfile {
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
	profile.SetAside = "unknown"
	if contract.SetAside != nil {
		profile.SetAside = *contract.SetAside
	}
	return profile
}
