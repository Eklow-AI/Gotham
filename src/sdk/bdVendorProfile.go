package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Eklow-AI/Gotham/src/models"
	"golang.org/x/oauth2"
)

var client *http.Client

func getRedShirtJWT() string {
	login := map[string]string{
		"username": "api_score",
		"password": "d627StTYf#y@lzg#Ej1*tmHL",
	}
	data, err := json.Marshal(login)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://redshirttest.g2xchange.com/wp-json/jwt-auth/v1/token",
		"application/json", bytes.NewBuffer(data))
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	var dataResp map[string]string
	json.NewDecoder(resp.Body).Decode(&dataResp)
	return dataResp["token"]
}

//SetupSDK initalizes the authorized client for sdk package
func SetupSDK() {
	token := getRedShirtJWT()
	ctx := context.Background()
	authorized := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
		TokenType:   "Bearer",
	}))
	client = authorized
}

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
		RecordPerPage: 5,
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
	for _, contract := range vendorHistory {
		label := contract[field].(string)
		_, exist := composition[label]
		if exist {
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

func main() {
	SetupSDK()
	vendorHistory := getVendorHistory("6ZP36")
	data := getComposition(vendorHistory, "funding_agency_name")
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))

}
