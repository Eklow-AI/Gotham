package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

var client *http.Client

func getRedShirtJWT() string {
	login := map[string]string{
		"username": "api_score",                //os.Getenv("rsUsername"),
		"password": "d627StTYf#y@lzg#Ej1*tmHL", //os.Getenv("rsPassword"),
	}
	data, err := json.Marshal(login)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://redshirttest.g2xchange.com/wp-json/jwt-auth/v1/token",
		"application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

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
