package ivtmodel

import (
	"fmt"
	"strconv"
)

const SystemDecimal = 1e8

type AmountCompareResult int8

const (
	Smaller AmountCompareResult = -1
	Equal   AmountCompareResult = 0
	Bigger  AmountCompareResult = 1
)

type InventoryAmount int64

func NewSystemAmount(value int64) InventoryAmount {
	return InventoryAmount(value)
}

func NewAmount[T Number](value T) InventoryAmount {
	return InventoryAmount(int64(float64(value) * SystemDecimal))
}

// ==== GETTER/SETTER FUNCTIONS ====

func (a InventoryAmount) SystemAmount() int64 {
	return int64(a)
}

func (a InventoryAmount) RealAmount() float64 {
	return float64(a) / SystemDecimal
}

func (a InventoryAmount) IntAmount() int {
	return int(a / SystemDecimal)
}

func (a InventoryAmount) String() string {
	floatVal := strconv.FormatFloat(a.RealAmount(), 'f', -1, 64)
	return fmt.Sprintf("%s (%d)", floatVal, a.SystemAmount())
}

// ==== MATHEMATICS FUNCTIONS ====

func (a InventoryAmount) Add(another InventoryAmount) InventoryAmount {
	return InventoryAmount(int64(a) + int64(another))
}

func (a InventoryAmount) Sub(another InventoryAmount) InventoryAmount {
	return InventoryAmount(int64(a) - int64(another))
}

func (a InventoryAmount) Mul(another InventoryAmount) InventoryAmount {
	return InventoryAmount(int64(a) * int64(another))
}

func (a InventoryAmount) Div(another InventoryAmount) InventoryAmount {
	return InventoryAmount(int64(a) / int64(another))
}

func (a InventoryAmount) Compare(another InventoryAmount) AmountCompareResult {
	switch {
	case a.SystemAmount() < another.SystemAmount():
		return Smaller
	case a.SystemAmount() > another.SystemAmount():
		return Bigger
	default:
		return Equal
	}
}
