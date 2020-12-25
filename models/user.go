package models

import "time"

// User struct defines the API user entity
type User struct {
	Email       string    `json:"email" gorm:"primary_key"`
	Name        string    `json:"name"`
	Token       string    `json:"token"`
	Utype       string    `json:"utype"`
	CallsToDate int       `json:"calls"`
	StartDate   time.Time `json:"start_date"`
	IsValid     bool      `json:"is_valid"`
}
