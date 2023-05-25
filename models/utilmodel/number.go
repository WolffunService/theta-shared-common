package utilmodel

import (
	"encoding/json"
	"math"
)

type Number struct {
	Value_  float64 `json:"value" mapstructure:"value" bson:"value" yaml:"value"` // dont use this directly, user n.Float() or n.Int()
	Decimal int     `json:"decimal" mapstructure:"decimal" bson:"decimal" yaml:"decimal"`
}

type numberJSON struct {
	Number `json:",inline"`
}

func NumberI(value float64) Number {
	return NumberF(value, 0)
}

func NumberDefault(value float64) Number {
	return NumberF(value, 8)
}

func NumberF(value float64, decimal int) Number {
	return Number{Value_: toFixed(value, decimal), Decimal: decimal}
}

func (n Number) ToInt() Number {
	n.Decimal = 0
	n.Value_ = float64(n.Int())
	return n
}

func (n *Number) SetInt(v int) *Number {
	*n = NumberI(float64(v))
	return n
}

// return float with fixed decimal points
func (n Number) Float() float64 {
	return toFixed(n.Value_, n.Decimal)
}

func (n Number) Real() float64 {
	return n.Value_
}

func (n Number) Int() int {
	return round(n.Value_)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func (n1 Number) Add(n2 Number) Number {
	return NumberF(n1.Value_+n2.Value_, n1.Decimal)
}

func (n1 Number) Sub(n2 Number) Number {
	return NumberF(n1.Value_-n2.Value_, int(math.Max(float64(n1.Decimal), float64(n2.Decimal))))
}

func (n1 Number) Mul(n2 Number) Number {
	return NumberF(n1.Value_*n2.Value_, int(math.Max(float64(n1.Decimal), float64(n2.Decimal))))
}

func (n1 Number) Div(n2 Number) Number {
	return NumberF(n1.Value_/n2.Value_, int(math.Max(float64(n1.Decimal), float64(n2.Decimal))))
}

func (n1 *Number) MarshalJSON() ([]byte, error) {
	n := numberJSON{Number: *n1}
	n.Value_ = n.Float()

	return json.Marshal(n)
}
