package currencymodel

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/WolffunService/thetan-shared-common/enums/currencyenum"
)

const SYSTEM_DECIMALS int = 8
const BLOCK_CHAIN_DECIMALS int = 18

type SystemCurrency struct {
	Type     currencyenum.Currency       `json:"type" bson:"type"`
	Name     string                      `json:"name,omitempty"  bson:"name,omitempty"`
	Value    int64                       `json:"value" bson:"value"`
	Decimals int                         `json:"decimals" bson:"decimals"`
	Source   currencyenum.CurrencySource `json:"-"  bson:"-"`
}

type SC struct {
	SystemCurrency `json:",inline" bson:",inline"`
}

func (s *SC) UnmarshalJSON(data []byte) error {
	if s == nil {
		return errors.New("cannot unmarshal nil pointer")
	}

	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	strs := strings.Split(str, "_")
	if len(strs) != 2 {
		return errors.New("cannot parse sc")
	}

	value, err := strconv.ParseFloat(strs[0], 64)
	if err != nil {
		return err
	}

	currencyType, err := strconv.Atoi(strs[1])
	if err != nil ||
		len(currencyenum.CurrencySymbol(currencyenum.Currency(currencyType))) == 0 {
		return err
	}

	s.SystemCurrency = ConvertFloatToSystemCurrency(value, currencyenum.Currency(currencyType))

	return nil
}

func (s SystemCurrency) ToAnalyticStr() string {
	return fmt.Sprintf("%d_%d_%d", s.Value, s.Type, s.Decimals)
}

type ExchangeRes struct {
	Success bool    `json:"success"`
	Data    float64 `json:"data"`
}

func (s SystemCurrency) GetRealValue() float64 {
	return float64(s.Value) / math.Pow10(SYSTEM_DECIMALS)
}

func (s SystemCurrency) Validate() bool {
	return true //TODO: Handle validate later //s.Value >= MAX_CURRENCY && s.Value <= MAX_CURRENCY_BLOCK_CHAIN && s.Decimals == SYSTEM_DECIMALS
}

func (s SystemCurrency) Compare(other SystemCurrency) currencyenum.CompareSystemCurrency {
	if s.Type != other.Type || s.Decimals != other.Decimals {
		return currencyenum.CS_DIFFERENCE
	}
	if s.Value < other.Value {
		return currencyenum.CS_SMALLER
	}
	if s.Value > other.Value {
		return currencyenum.CS_BIGGER
	}
	return currencyenum.CS_EQUAL
}

func (s SystemCurrency) Multiple(mul int64) SystemCurrency {
	return SystemCurrency{
		Type:     s.Type,
		Name:     s.Name,
		Value:    s.Value * mul,
		Decimals: s.Decimals,
	}
}

func (s SystemCurrency) MultipleFloat(mul float64) SystemCurrency {
	return SystemCurrency{
		Type:     s.Type,
		Name:     s.Name,
		Value:    int64(float64(s.Value) * mul), //TODO: Anh Lam check lai giup em
		Decimals: s.Decimals,
	}
}

func ConvertIntToSystemCurrency(realValue int, typeCurrency currencyenum.Currency) SystemCurrency {
	//TODO: validate realValue
	systemValue := int64(math.Pow10(SYSTEM_DECIMALS)) * int64(realValue)
	return convertSystemCurrency(systemValue, typeCurrency)
}

// khi lay price da duoc convert roi
func ConvertInt64ToSystemCurrencyWithSystemValue(realValue int64, typeCurrency currencyenum.Currency) SystemCurrency {
	//TODO: validate realValue
	return convertSystemCurrency(realValue, typeCurrency)
}

func ConvertInt64ToSystemCurrency(realValue int64, typeCurrency currencyenum.Currency) SystemCurrency {
	//TODO: validate realValue
	systemValue := int64(math.Pow10(SYSTEM_DECIMALS)) * realValue
	return convertSystemCurrency(systemValue, typeCurrency)
}

func ConvertFloatToSystemCurrency(realValue float64, typeCurrency currencyenum.Currency) SystemCurrency {
	//TODO: validate realValue
	systemValue := int64(math.Round(math.Pow10(SYSTEM_DECIMALS) * realValue))
	return convertSystemCurrency(systemValue, typeCurrency)
}

func ConvertFromAmountInWieBlockChain(valueInBlockChain string, typeCurrency currencyenum.Currency) (*SystemCurrency, bool) {

	amountInWei := new(big.Int)
	amountInWei, ok := amountInWei.SetString(valueInBlockChain, 10)
	if !ok {
		log.Printf("Amount in wie error, amountInWei = %v", valueInBlockChain)
		return nil, false
	}
	systemValue := new(big.Int).Div(amountInWei, big.NewInt(int64(math.Pow10(BLOCK_CHAIN_DECIMALS-SYSTEM_DECIMALS))))
	systemCurrency := &SystemCurrency{
		Type:     typeCurrency,
		Name:     currencyenum.CurrencySymbol(typeCurrency),
		Value:    systemValue.Int64(),
		Decimals: SYSTEM_DECIMALS,
	}
	return systemCurrency, true
}

func (s SystemCurrency) ConvertValueToWieBlockChain() *big.Int {
	systemValue := big.NewInt(s.Value)
	return new(big.Int).Mul(systemValue, big.NewInt(int64(math.Pow10(BLOCK_CHAIN_DECIMALS-SYSTEM_DECIMALS))))
}

func convertSystemCurrency(systemValue int64, typeCurrency currencyenum.Currency) SystemCurrency {
	systemCurrency := SystemCurrency{
		Type:     typeCurrency,
		Name:     currencyenum.CurrencySymbol(typeCurrency),
		Value:    systemValue,
		Decimals: SYSTEM_DECIMALS,
	}
	return systemCurrency
}

func NewSystemCurrency(systemValue int64, typeCurrency currencyenum.Currency) SystemCurrency {
	return convertSystemCurrency(systemValue, typeCurrency)
}
