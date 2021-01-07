package models

import "testing"

func TestUtypeRankings(t *testing.T) {
	if utypeRanking["non"] != 0 {
		t.Errorf("broken utype ranking: `non` should be 0")
	}
	if utypeRanking["trial"] != 1 {
		t.Errorf("broken utype ranking: `trial` should be 1")
	}
	if utypeRanking["pro"] != 2 {
		t.Errorf("broken utype ranking: `pro` should be 2")
	}
	if utypeRanking["enterprise"] != 3 {
		t.Errorf("broken utype ranking: `enterprise` should be 3")
	}
}
