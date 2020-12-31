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
func CreateToken(email string) (token string, err error) {
	if email == "" {
		return "", errors.New("Token Services: Email cannot be empty")
	}
	hashedEmail := hash(email)
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	token = fmt.Sprintf("%d.%s", hashedEmail, uuid)
	return token, nil
}
