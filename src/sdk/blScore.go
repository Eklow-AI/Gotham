package sdk

import (
	"math"
	"sort"
)

// GetScore returns the compatability score of a vendor with a contract
func GetScore(cage string, cid string) (score float64) {
	vendor := getVendorProfile(cage)
	contract := getContract(cid)
	coScore, w0 := calcCOScore(contract.CO, vendor.COs)
	acScore, w1 := calcContractAgencyScore(contract.ContractAgency, vendor.ContractAgency)
	naicsScore, w2 := calcNaics(contract.Naics, vendor.Naics)
	pscScore, w3 := calcPsc(contract.Psc, vendor.Psc)
	setAsideScore, w4 := calcSetAside(contract.SetAside, vendor.SetAsides)
	coSizeScore, w5 := calcCOSize(contract.COSizeSelection, vendor.COSizeSelection)
	offerScore, numOffers := 0.0, float64(contract.NumOffers)
	if numOffers == 0 {
		offerScore = 1
	} else {
		offerScore = capScore(2*math.Log(numOffers)+0.5, 10)
	}
	totalWeight := w0 + w1 + w2 + w3 + w4 + w5
	score = coScore + acScore + naicsScore + pscScore + offerScore + setAsideScore + coSizeScore
	score = score / totalWeight
	return score
}

func calcCOScore(co string, COs map[string]float64) (score float64, weight float64) {
	weight = 25.0
	floor := weight * 0.05
	type sortedCO struct {
		name        string
		probability float64
	}

	if co == "unkown" || len(COs) == 0 {
		return floor, weight
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
	return score, weight
}

func calcContractAgencyScore(agency string, agencies map[string]float64) (score float64, weight float64) {
	weight = 20.0
	floor := 0.25 * weight
	if _, ok := agencies[agency]; !ok {
		return floor, weight
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
	return score, weight
}

func calcSetAside(setAside string, setAsides map[string]float64) (score float64, weight float64) {
	weight = 30.0
	floor := weight * 0.18
	// Check if set-aside has been used before by the company
	if _, ok := setAsides[setAside]; !ok {
		return floor, weight
	}
	// Deal with people who usually don't have set-asides
	if setAside == "unknown" && setAsides[setAside] > 0.45 {
		return weight * 0.55, weight
	}
	type sortedSA struct {
		name        string
		probability float64
	}
	// rank Agencies based on their freq give a score
	probsArr := make([]sortedSA, len(setAsides))
	idx := 0
	for key, value := range setAsides {
		probsArr[idx] = sortedSA{name: key, probability: value}
		idx++
	}
	sort.Slice(probsArr, func(i, j int) bool {
		return probsArr[i].probability > probsArr[j].probability
	})
	probLen := float64(len(setAsides))
	topBracket := int(probLen*0.667 + 0.5)
	midBracket := int(probLen*0.333 + 0.5)
	for idx, value := range probsArr {
		switch {
		case value.name == setAside && idx < midBracket:
			score = weight*value.probability + floor + 1
		case value.name == setAside && idx >= midBracket && idx < topBracket:
			score = weight * 0.8
		case value.name == setAside && idx >= topBracket:
			score = weight
		default:
			continue
		}
	}
	return score, weight
}

func calcNaics(naics string, naicsDist map[string]float64) (score float64, weight float64) {
	weight = 15.0
	floor := weight * 0.1
	if _, ok := naicsDist[naics]; !ok {
		return floor, weight
	} else if naicsDist[naics] < 0.5 {
		score = (naicsDist[naics]*weight + 0.02) + floor
	} else {
		score = weight
	}
	return score, weight
}

func calcPsc(psc string, pscDist map[string]float64) (score float64, weight float64) {
	weight = 15.0
	if _, ok := pscDist[psc]; !ok {
		return 0, weight
	} else if pscDist[psc] < 0.5 {
		score = weight * (pscDist[psc] + 0.05)
	} else {
		score = weight
	}
	return score, weight
}

func calcCOSize(size string, sizes map[string]float64) (score float64, weight float64) {
	// TODO call XFactor API to check if vendor is small biz
	weight = 25.0
	// Check if the size selection is small
	// but vendor is not small, then return floor
	_, isSmall := sizes["SMALL BUSINESS"]
	if size == "SMALL BUSINESS" && !isSmall {
		return weight * 0.20, weight
	}
	// If there is no size selection
	//
	if size != "SMALL BUSINESS" {
		return 0, 0
	}
	if isSmall && size == "SMALL BUSINESS" {
		probability := sizes[size]
		switch {
		case sizes[size] < 0.5:
			score = weight * (probability + 0.225)
		default:
			score = weight
		}
	}
	return score, weight
}

func capScore(score, cap float64) float64 {
	if score > cap {
		return cap
	}
	return score
}
