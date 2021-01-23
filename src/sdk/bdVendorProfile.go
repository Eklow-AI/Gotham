package sdk

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"

	"github.com/Eklow-AI/Gotham/src/models"
)

func getVendorHistory(cage string) (vendorHistory []map[string]interface{}) {
	query := models.RedShirtQuery{
		Object:      "contracts",
		Version:     "1.0",
		Timeout:     45000,
		RecordLimit: 10000,
		Rows:        true,
		Totals:      true,
		Lists:       false,
		SearchFilter: []models.Filter{
			{
				Field:    "cage_code",
				Operator: "tsquery",
				Value:    cage,
			},
		},
		RecordPerPage: 120,
		CurrentPage:   1,
		SortFilter: []models.Filter{
			{
				Field: "date_signed",
				Order: "desc",
			},
		},
	}
	data, err := json.Marshal(query)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Post("https://redshirttest.g2xchange.com/wp-json/api/v1/query/?db=MM10TEST",
		"application/json", bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	var dataResp models.RedShirtResp
	json.NewDecoder(resp.Body).Decode(&dataResp)
	vendorHistory = dataResp.ListData
	return vendorHistory
}

func getComposition(vendorHistory []map[string]interface{}, field string) (composition map[string]float64) {
	composition = make(map[string]float64)
	for _, contract := range vendorHistory {
		label, ok := contract[field].(string)
		if !ok {
			label = "unknown"
		}
		if _, ok := composition[label]; ok {
			composition[label] += 1.0
		} else {
			composition[label] = 1.0
		}
	}
	totalContracts := float64(len(vendorHistory))
	for key, value := range composition {
		composition[key] = value / totalContracts
	}
	return composition
}

func getVendorProfile(cage string) (profile models.VendorProfile) {
	vendorHistory := getVendorHistory(cage)
	profile.Name = vendorHistory[0]["vendor_name"].(string)
	profile.Cage = cage
	profile.FundingAgency = getComposition(vendorHistory, "funding_agency_name")
	profile.Naics = getComposition(vendorHistory, "naics_code")
	profile.Psc = getComposition(vendorHistory, "product_or_service_code_text")
	profile.SetAsides = getComposition(vendorHistory, "type_of_set_aside_description")
	profile.COSizeSelection = getComposition(vendorHistory, "contracting_officer_business_size_determination_description")
	profile.COs = getComposition(vendorHistory, "last_modified_by")
	profile.PlacesOfPerf = getComposition(vendorHistory, "place_of_performance_zip_code_city")
	// Cast string to int because redshirt returns Zip codes as strings
	sZip := vendorHistory[0]["phy_zip_code"].(string)
	zip, err := strconv.ParseInt(sZip, 10, 32)

	if err != nil {
		log.Fatal(err)
	}

	profile.Zip = zip

	return profile
}
