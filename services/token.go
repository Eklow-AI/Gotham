package services

import (
	"errors"
	"fmt"
	"strings"

	"hash/fnv"

	"github.com/google/uuid"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// CreateToken returns an API token string given user email, patron, and security level
func CreateToken(email string, patron string, utype string) (token string, err error) {
	if email == "" {
		return "", errors.New("Token Services: Email cannot be empty")
	}
	if patron == "" {
		return "", errors.New("Token Services: Patron cannot be empty")
	}
	if utype == "" {
		return "", errors.New("Token Services: Security level cannot be 0")
	}
	hashedEmail := hash(email)
	hashedPatron := hash(patron)
	hashedUtype := hash(utype)
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	token = fmt.Sprintf("%d.%s.%d.%d", hashedEmail, uuid, hashedPatron, hashedUtype)
	return token, nil
}
