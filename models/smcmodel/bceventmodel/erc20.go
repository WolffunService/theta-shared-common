package bceventmodel

type Deposit struct {
	Block           uint64
	TransactionHash string
	From            string
	PaymentToken    string
	AmountInWei     string
}
