package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Eklow-AI/Gotham/src/api"
	"github.com/jinzhu/gorm"
)

// Credentials struct used to create a new login
type Credentials struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email" binding:"Required"`
	Password  string `json:"password" binding:"Required"`
}

// CreateUser http request function User API handler
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `Credentials` instance
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// if there's something wrong with the body, return a 400 request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = api.CreateUser(creds.FirstName, creds.LastName, creds.Email, creds.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := map[string]bool{"success": true}
	json.NewEncoder(w).Encode(resp)
}

// SignInUser http request function User API handler
func SignInUser(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `Credentials` instance
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// if there's something wrong with the body, return a 400 request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	cookie, err := api.SignIn(creds.Email, creds.Password)
	// Error handling
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		resp := map[string]interface{}{"success": false, "error": "User does not exist"}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(resp)
		return
	} else if err != nil && cookie == nil {
		resp := map[string]interface{}{"success": false, "error": "Incorrect username or password"}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(resp)
		return
	} else if err != nil {
		resp := map[string]interface{}{"success": false, "error": err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}
	// if everything goes well, then we send the cookie :D
	http.SetCookie(w, cookie)
	resp := map[string]bool{"success": true}
	json.NewEncoder(w).Encode(resp)
}
