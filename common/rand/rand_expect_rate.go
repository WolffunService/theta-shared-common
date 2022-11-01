package rand

import (
	"errors"
	"math"
)

var ErrEmptyExpectRate = errors.New("empty expect rate")
var ErrRateOutOfRange = errors.New("rate must in range 0.0 -> 1.0")

func PickItemByExpectRate[K comparable](appearedItem map[K]int, expectRate map[K]float64) (K, error) {
	if len(expectRate) == 0 {
		return *new(K), ErrEmptyExpectRate
	}

	totalAppeared := 0
	tempAppeared := make(map[K]int, len(appearedItem))
	for item, nAppeared := range appearedItem {
		tempAppeared[item] = nAppeared
		totalAppeared += nAppeared
	}

	// Calculate (Actual Rate - Expect Rate)
	currentRate := make(map[K]float64)
	var minRate = math.MaxFloat64
	var itemMinRate K
	var actualRate float64
	for item, rate := range expectRate {
		if rate < 0.0 || rate > 1.0 {
			return *new(K), ErrRateOutOfRange
		}

		actualRate = 0.0
		if totalAppeared != 0 {
			actualRate = float64(tempAppeared[item]) / float64(totalAppeared)
		}
		currentRate[item] = actualRate - rate

		if currentRate[item] < minRate {
			minRate = currentRate[item]
			// Copy memory in foreach
			pickItem := item
			itemMinRate = pickItem
		}
	}

	return itemMinRate, nil
}
