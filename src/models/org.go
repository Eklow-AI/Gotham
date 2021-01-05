package models

import (
	"errors"
	"fmt"

	"github.com/Eklow-AI/Gotham/src/services"
)

//Org defines orgs that are allowed to use the Gotham API
//Only other orgs with a security Clearance of 3 can create other orgs
type Org struct {
	Email               string `json:"email"`
	Name                string `json:"name"`
	Utype               string `json:"utype"`
	Clearance           int    `json:"clearance"`
	Token               string `gorm:"primary_key"`
	NumUsers            int    `gorm:"default:0"`
	ActiveSubscriptions int    `gorm:"default:0"`
	IsValid             bool
}

// NewOrgOptions is a struct for specific configuration options for a user
type NewOrgOptions struct {
	Email     string `json:"email" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Utype     string `json:"utype" binding:"required"`
	Clearance int    `json:"clearance" binding:"required"`
}

//InsertNewOrg inserts a new user struct to the db
func InsertNewOrg(options NewOrgOptions) (org *Org, err error) {
	token, err := services.CreateToken(options.Email)
	if err != nil {
		return &Org{}, errors.New(fmt.Sprintln("error inserting org:", err))
	}
	org = &Org{
		Email:     options.Email,
		Name:      options.Name,
		Utype:     options.Utype,
		Clearance: options.Clearance,
		Token:     token,
		IsValid:   true,
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
