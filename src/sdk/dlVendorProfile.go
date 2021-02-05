package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/Eklow-AI/Gotham/src/models"
)

func getVendorHistory(cage string) []models.ContractData {
	query := models.RedShirtQuery{
		Object:      "contracts",
		Version:     "1.2",
		Timeout:     45000,
		RecordLimit: 10000,
		Rows:        true,
		Totals:      false,
		Lists:       false,
		SearchFilter: []models.Filter{
			{
				Field:    "cage_code",
				Operator: "tsquery",
				Value:    cage,
			},
		},
		RecordPerPage: 60,
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
	contractHistory := dataResp.ContractHistory
	return contractHistory
}

func getComposition(vendorHistory []models.ContractData, field string) map[string]float64 {
	composition := make(map[string]float64)
	for _, contract := range vendorHistory {
		reflection := reflect.ValueOf(contract)
		rValue := reflect.Indirect(reflection).FieldByName(field)
		label := rValue.String()
		// Check that the reflected string is valid
		if label == "<invalid Value>" || strings.Contains(label, " Value>") {
			log.Fatal("Problem calculating composition; reflected string is not valid")
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

func getVendorProfile(cage string) models.VendorProfile {
	vendorHistory := getVendorHistory(cage)
	profile := models.VendorProfile{}
	profile.Cage = cage
	profile.Name = vendorHistory[0].VendorName
	profile.ContractAgency = getComposition(vendorHistory, "ContractAgency")
	profile.Naics = getComposition(vendorHistory, "Naics")
	profile.Psc = getComposition(vendorHistory, "Psc")
	profile.SetAsides = getComposition(vendorHistory, "SetAside")
	profile.COSizeSelection = getComposition(vendorHistory, "COSizeSelection")
	profile.COs = getComposition(vendorHistory, "CO")
	profile.PlacesOfPerf = getComposition(vendorHistory, "PlaceOfPerf")
	// Cast string to int because redshirt returns Zip codes as strings
	profile.Zip = vendorHistory[0].VendorZip
	return profile
}

func getCVendorProfile(cage string, c chan models.VendorProfile) {
	fmt.Println("A")
	c <- getVendorProfile(cage)
	fmt.Println("A")
}
