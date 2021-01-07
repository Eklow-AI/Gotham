package services

import (
	"testing"
)

func TestUtypeRankings(t *testing.T) {
	token := CreateToken("biggie@smallz.com")
	if len(token) < 35 {
		t.Errorf("Token length is to short, should be greater than 35 characters")
	}
}
