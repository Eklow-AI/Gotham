package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User defines the user object struct
type User struct {
	UUID           uuid.UUID `json:"uuid" gorm:"type:uuid;primary_key"`
	Email          string    `json:"email" form:"email" gorm:"unique"`
	Password 	   string    `json:"password"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}

// Claims defines the users's JWT claim model
type Claims struct {
	Email string `json:"username"`
	jwt.StandardClaims
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("UUID", uuid.NewV4())
}
