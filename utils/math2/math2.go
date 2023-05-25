package math2

import (
	crand "crypto/rand"
	"math"
	"math/big"
)

// RandomInt return a number from min to max - 1
func RandomInt(min, max int) (int, error) {
	i, err := Random0ToInt(max - min)
	if err != nil {
		return max, nil
	}
	i += min
	return i, nil
}

// Random0ToInt return a number from 0 to max - 1, return 0 if max == 0 and return error if max's negative
func Random0ToInt(max int) (int, error) {
	if max == 0 {
		return 0, nil
	}
	preRand, err := crand.Int(crand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return -1, err
	}
	return int(preRand.Int64()), nil
}

// Random0ToInt return a number from 0 to max - 1, return 0 if max == 0 and return error if max's negative
func Random0ToInt64(max int64) (int64, error) {
	if max == 0 {
		return 0, nil
	}
	preRand, err := crand.Int(crand.Reader, big.NewInt(max))
	if err != nil {
		return -1, err
	}
	return preRand.Int64(), nil
}

func AbsInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Maxf(x, y float64) float64 {
	return math.Max(x, y)
}

func Minf(x, y float64) float64 {
	return math.Min(x, y)
}

//// DEPRECATED:  the feature is marked "deprecated" and should no longer be used. Use math.Round2 instead.
//// EX: Round(float64(100) / float64(7), 0.00001) --> 14.285710000000002
//func Round(x, unit float64) float64 {
//	return math.Round(x/unit) * unit
//}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// EX: Round2(float64(100) / float64(7), 100000) --> 14.28571
func Round2(x, unit float64) float64 {
	return math.Round(x*unit) / unit
}

func IntInRange(i, min, max int) int {
	switch {
	case i < min:
		return min
	case i > max:
		return max
	default:
		return i
	}
}

// EX: Ceil2(float64(100) / float64(7), 100000)
func Ceil2(x, unit float64) float64 {
	return math.Ceil(x*unit) / unit
}

// EX: Floor2(float64(100) / float64(7), 100000)
func Floor2(x, unit float64) float64 {
	return math.Floor(x*unit) / unit
}

func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}
