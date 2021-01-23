package models

import (
	"errors"
	"fmt"
)

//Org defines orgs that are allowed to use the Gotham API
//Only other orgs with a security Clearance of 3 can create other orgs
type Org struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	CallsToDate int    `gorm:"default:0"`
}

// NewOrgOptions is a struct for specific configuration options for a user
type NewOrgOptions struct {
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Utype string `json:"utype" binding:"required"`
}

//InsertNewOrg inserts a new user struct to the db
func InsertNewOrg(options NewOrgOptions) (org *Org, err error) {
	org = &Org{
		Email: options.Email,
		Name:  options.Name,
	}
	result := DB.Create(org)
	if result.Error != nil {
		return &Org{}, errors.New(fmt.Sprintln("error inserting org:", result.Error))
	}
	return org, nil
}

// GetOrg gets an org from the token (pk)
func GetOrg(token string) (org *Org) {
	org = &Org{}
	DB.Where("token = ?", token).First(org)
	return org
}
