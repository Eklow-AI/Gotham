package models

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// User struct defines the API user entity
type User struct {
	Email       string    `json:"email" gorm:"primary_key"`
	Name        string    `json:"name"`
	Utype       string    `json:"utype"`
	CallsToDate int       `json:"calls" gorm:"default:0"`
	Created     time.Time `json:"start_date" gorm:"default:'current_timestamp'"`
	IsValidDate time.Time `json:"is_valid_date"`
}

// NewUserOptions is a struct for specific configuration options for a user
type NewUserOptions struct {
	Email       string `json:"email" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Utype       string `json:"utype" binding:"required"`
	IsValidDate time.Time
}

//define hierachy of user types
var utypeRanking = map[string]int{
	"non":        0,
	"trial":      1,
	"pro":        2,
	"enterprise": 3,
}

// GetUser returns a pointer to a user struct giver an email (pk)
func getUser(email string) (user *User) {
	user = &User{}
	DB.Where("email = ?", email).First(user)
	return user
}

//InsertNewUser inserts a new user struct to the db
func InsertNewUser(options NewUserOptions) (err error) {
	if strings.ToLower(options.Utype) == "trial" {
		// Give a trialing user a two week trial
		options.IsValidDate = time.Now().AddDate(0, 0, 7*2)
	} else {
		// Give other users a 1 year subscription
		options.IsValidDate = time.Now().AddDate(1, 0, 0)
	}
	user := &User{
		Email:       options.Email,
		Name:        options.Name,
		Utype:       options.Utype,
		IsValidDate: options.IsValidDate,
	}
	result := DB.Create(user)
	if result.Error != nil {
		return errors.New(fmt.Sprintln("error inserting user:", result.Error))
	}
	return nil
}

// SetUtype sets the utype of user and updates it on the database
func (user *User) SetUtype(utype string) (err error) {
	if utype == "" {
		return errors.New("utype cannot be empty")
	}
	// Check if it is a valid utype
	if _, isThere := utypeRanking[utype]; !isThere {
		return errors.New("utype not valid")
	}
	// this means a trialing user is being upgraded
	if utypeRanking[user.Utype] == 1 && utypeRanking[utype] > 1 {
		DB.Model(user).Update("IsValidDate", time.Now().AddDate(1, 0, 0))
	} else if utypeRanking[utype] == 0 {
		// user is being removed from memebership
		DB.Model(user).Update("IsValidDate", time.Now())
	}
	DB.Model(user).Update("utype", utype)
	user.Utype = utype
	return nil
}
