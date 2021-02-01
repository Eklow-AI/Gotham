package sdk

import (
	"math"
	"sort"
)

// GetScore returns the compatability score of a vendor with a contract
func GetScore(cage string, cid string) (score float64) {
	vendor := getVendorProfile(cage)
	contract := getContract(cid)
	// Get CO experience scores
	coScore := calcCOScore(contract.CO, vendor.COs)
	acScore := calcContractAgencyScore(contract.ContractAgency, vendor.ContractAgency)
	naicsScore := calcNaics(contract.Naics, vendor.Naics)
	pscScore := calcPsc(contract.Psc, vendor.Psc)
	setAsideScore := calcSetAside(contract.SetAside, vendor.SetAsides)

	offerScore, numOffers := 0.0, float64(contract.NumOffers)
	if numOffers == 0 {
		offerScore = 1
	} else {
		offerScore = capScore(2*math.Log(numOffers)+0.5, 10)
	}

	score = coScore + acScore + naicsScore + pscScore + offerScore + setAsideScore
	return score
}

func calcCOScore(co string, COs map[string]float64) (score float64) {
	weight := 25.0
	floor := weight * 0.05
	type sortedCO struct {
		name        string
		probability float64
	}

	if co == "unkown" || len(COs) == 0 {
		return floor
	}
	// rank COs and based on their freq give a score
	probsArr := make([]sortedCO, len(COs))
	idx := 0
	for key, value := range COs {
		probsArr[idx] = sortedCO{name: key, probability: value}
		idx++
	}
	sort.Slice(probsArr, func(i, j int) bool {
		return probsArr[i].probability > probsArr[j].probability
	})
	probLen := float64(len(probsArr))
	topBracket := int(probLen*0.666+0.5) - 1
	midBracket := int(probLen*0.333+0.5) - 1
	for idx, value := range probsArr {
		switch {
		case value.name == co && idx < midBracket:
			score = weight*value.probability + floor
		case value.name == co && idx >= midBracket && idx < topBracket:
			score = weight*value.probability + (2.0 * floor)
		case value.name == co && idx >= topBracket:
			score = weight
		default:
			continue
		}
	}
	return score
}

func calcContractAgencyScore(agency string, agencies map[string]float64) (score float64) {
	weight := 20.0
	floor := 0.25 * weight
	if _, ok := agencies[agency]; !ok {
		return floor
	}
	type sortedAgency struct {
		name        string
		probability float64
	}

	// rank Agencies based on their freq give a score
	probsArr := make([]sortedAgency, len(agencies))
	idx := 0
	for key, value := range agencies {
		probsArr[idx] = sortedAgency{name: key, probability: value}
		idx++
	}
	sort.Slice(probsArr, func(i, j int) bool {
		return probsArr[i].probability > probsArr[j].probability
	})
	probLen := float64(len(agencies))
	topBracket := int(probLen*0.667 + 0.5)
	midBracket := int(probLen*0.333 + 0.5)
	for idx, value := range probsArr {
		switch {
		case value.name == agency && idx < midBracket:
			score = (1.0 + value.probability) * floor
		case value.name == agency && idx >= midBracket && idx < topBracket:
			score = weight*value.probability + floor
		case value.name == agency && idx >= topBracket:
			score = weight
		default:
			continue
		}
	}
	return score
}

func calcSetAside(setAside string, setAsides map[string]float64) (score float64) {
	weight := 30.0
	floor := weight * 0.135
	// Check if set-aside has been used before by the company
	if _, ok := setAsides[setAside]; !ok {
		return floor
	}
	

	return score
}

func calcNaics(naics string, naicsDist map[string]float64) (score float64) {
	weight := 15.0
	floor := weight * 0.1
	if _, ok := naicsDist[naics]; !ok {
		return floor
	} else if naicsDist[naics] < 0.5 {
		score = (naicsDist[naics]*weight + 0.02) + floor
	} else {
		score = weight
	}
	return score
}

func calcPsc(psc string, pscDist map[string]float64) (score float64) {
	weight := 15.0
	if _, ok := pscDist[psc]; !ok {
		return 0
	} else if pscDist[psc] < 0.5 {
		score = weight * (pscDist[psc] + 0.05)
	} else {
		score = weight
	}
	return score
}

func capScore(score, cap float64) float64 {
	if score > cap {
		return cap
	}
	return score
}
