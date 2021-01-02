package models

import (
	"errors"
	"fmt"
	"time"
)

// User struct defines the API user entity
type User struct {
	Email       string    `json:"email" gorm:"primary_key"`
	Name        string    `json:"name"`
	Utype       string    `json:"utype"`
	CallsToDate int       `json:"calls"`
	Created     time.Time `json:"start_date"`
	IsValid     bool      `json:"is_valid"`
}

// NewUserOptions is a struct for specific configuration options for a user
type NewUserOptions struct {
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Utype string `json:"utype" binding:"required"`
}

//InsertNewUser inserts a new user struct to the db
func InsertNewUser(options NewUserOptions) (err error) {
	user := &User{
		Email:       options.Email,
		Name:        options.Name,
		Utype:       options.Utype,
		CallsToDate: 0,
		Created:     time.Now(),
		IsValid:     true,
	}
	result := DB.Create(user)
	if result.Error != nil {
		return errors.New(fmt.Sprintln("error inserting user:", result.Error))
	}
	return nil
}
