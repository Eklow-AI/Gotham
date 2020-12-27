package models

import (
	"errors"
	"time"

	"github.com/Eklow-AI/Gotham/services"
)

// User struct defines the API user entity
type User struct {
	Email       string    `json:"email" gorm:"primary_key"`
	Name        string    `json:"name"`
	Token       string    `json:"token"`
	Utype       string    `json:"utype"`
	CallsToDate int       `json:"calls"`
	StartDate   time.Time `json:"start_date"`
	IsValid     bool      `json:"is_valid"`
	patron      string    `json:"patron"`
}

// Options is a struct for specific configuration options for a user
type Options struct {
	Email  string
	Name   string
	Utype  string
	Patron string
}

// New returns an instance of User. All newly created users are considered valid users
func New(options Options) (user User, err error) {

	if options.Email == "" {
		return User{}, errors.New("User Models: Email cannot be empty")
	}
	if options.Name == "" {
		return User{}, errors.New("User Models: Name cannot be empty")
	}
	if options.Utype == "" {
		return User{}, errors.New("User Models: Utype cannot be empty")
	}
	token, err := services.CreateToken(options.Email, options.Patron, options.Email)
	if err != nil {
		return User{}, err
	}
	return User{
		Email:       options.Email,
		Name:        options.Name,
		Token:       token,
		Utype:       options.Utype,
		CallsToDate: 0,
		patron:      options.Patron,
		StartDate:   time.Now(),
		IsValid:     true,
	}, nil
}
