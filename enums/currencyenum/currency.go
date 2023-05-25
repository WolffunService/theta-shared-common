package currencyenum

import "fmt"

var IsMainNet bool

// Currency
type Currency int // DUPLICATE VOI THETA-REPORT, sua thi sua ca hai

const (
	GTHC Currency = 1
	GTHG Currency = 2
	GPP  Currency = 3
	GMGT Currency = 4
	GPT  Currency = 5
	GMQP Currency = 6

	THC Currency = 11
	THG Currency = 12

	BUSD Currency = 31
	WBNB Currency = 32
	BNB  Currency = 33

	BOXC Currency = 1001
	BOXE Currency = 1002
	BOXL Currency = 1003
	BOXS Currency = 1004
)

var currencySymbol = map[Currency]string{
	GTHC: "gTHC",
	GTHG: "gTHG",
	GPP:  "gPP",
	GMGT: "gMGT",
	GPT:  "gQT",
	GMQP: "gMQP",
	THC:  "THC",
	THG:  "THG",
	BUSD: "BUSD",
	WBNB: "WBNB",
}

var currencyName = map[Currency]string{
	GTHC: "Thetan Coin In-game",
	GTHG: "Thetan Gem In-game",
	GPP:  "Power Point",
	GMGT: "Mini Game Ticket",
	GPT:  "Private Ticket",
	GMQP: "Marketplace Quest Point",

	THC: "Thetan Coin",
	THG: "Thetan Game",

	BUSD: "Binance USD",
	WBNB: "Wrapped BNB",
}

var currencyPaymentContract = map[Currency]string{
	GTHC: "",
	GTHG: "",
	GPP:  "",
	GMGT: "",
	GPT:  "",
	GMQP: "",
	THC:  "0x21f9b2137d4e1b83d09ea373be773a986c0e3f90",
	THG:  "0x948deddc3d3cf03017fd4e42cc9ac874402d0bb9",
	BUSD: "0x1d7256253f06cd8e1740ba0edda4d70b1a852b08",
	WBNB: "0x15c9e651b5971feb66e19fe9e897be6bdc3e841a",
}

var currencyPaymentContractMainNet = map[Currency]string{
	GTHC: "",
	GTHG: "",
	GPP:  "",
	GMGT: "",
	GPT:  "",
	GMQP: "",
	THC:  "0x24802247bd157d771b7effa205237d8e9269ba8a",
	THG:  "0x9fd87aefe02441b123c3c32466cd9db4c578618f",
	BUSD: "0xe9e7cea3dedca5984780bafc599bd69add087d56",
	WBNB: "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c",
}

func CurrencySymbol(code Currency) string {
	return currencySymbol[code]
}

func (currency Currency) Symbol() string {
	return CurrencySymbol(currency)
}

func CurrencyName(currency Currency) string {
	if name, found := currencyName[currency]; found {
		return name
	}

	return fmt.Sprintf("Unknown currency (%d)", int(currency))
}

func (currency Currency) Name() string {
	return CurrencyName(currency)
}

func CurrencyPaymentContract(code Currency) string {
	if IsMainNet {
		return currencyPaymentContractMainNet[code]
	} else {
		return currencyPaymentContract[code]
	}
}

func (currency Currency) PaymentContract() string {
	return CurrencyPaymentContract(currency)
}

//func GetPaymentContractCashOut(code Currency) (string, error) {
//	if code == GTHC {
//		return CurrencyPaymentContract(THC), nil
//	} else if code == GTHG {
//		return CurrencyPaymentContract(THG), nil
//	}
//
//	return "", fmt.Errorf("Do not support claim this currency %v", code)
//}

type TransactionStatus int

const (
	TS_NONE          TransactionStatus = 0
	TS_NEW           TransactionStatus = 1
	TS_PENDING       TransactionStatus = 2
	TS_SUCCESS       TransactionStatus = 3
	TS_WAITING_ADMIN TransactionStatus = 4
	TS_ADMIN_REVERT  TransactionStatus = 10
	TS_ADMIN_REJECT  TransactionStatus = 97 //send qua rabbitmq bi loi
	TS_SEND_ERROR    TransactionStatus = 98 //send qua rabbitmq bi loi
	TS_ERROR         TransactionStatus = 99
	TS_REFUND        TransactionStatus = 100
)

func (t TransactionStatus) IsWaiting() bool {
	return t == TS_NEW || t == TS_PENDING || t == TS_WAITING_ADMIN
}

type TransactionType int

const (
	DEPOSIT TransactionType = 1
	CLAIM   TransactionType = 2
)

type CompareSystemCurrency int

const (
	CS_SMALLER    CompareSystemCurrency = -1
	CS_EQUAL      CompareSystemCurrency = 0
	CS_BIGGER     CompareSystemCurrency = 1
	CS_DIFFERENCE CompareSystemCurrency = 99
)
