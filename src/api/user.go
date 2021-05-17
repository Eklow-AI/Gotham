package api

import (
	"net/http"
	"os"
	"time"

	"github.com/Eklow-AI/Gotham/src/models"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("jwtKey"))

// CreateUser salts and hashes user password using bycript and
// create a new user object in postgres
func CreateUser(firstName, lastName, email, password string) error {
	// Hash the user password
	HashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	// Create a user model instance and the add it to the postgres database
	user := models.User{
		Email:    email,
		Password: string(HashedPass[:]),
	}
	result := models.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SignIn checks user sign in info and returns a boolean whether it
// is a valid log in
func SignIn(email, password string) (cookie *http.Cookie, err error) {
	// Get the user from the database and its credentials
	var user models.User
	result := models.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	// Compare the stored hashed password, with the hashed version of the password that was received
	// if the passwords do not match return isValid false
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, nil
	}
	// Declare expiration time for the JWT, here we keep it at 72 hours
	expirationDate := time.Now().Add(72 * time.Hour)
	// Establish claims that will go inside the JWT
	claims := &models.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationDate.Unix(),
		},
	}
	// Create the JWT and then sign it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}
	// Create an cookie and attach the token to it
	cookie = &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationDate,
	}
	return cookie, nil
}
