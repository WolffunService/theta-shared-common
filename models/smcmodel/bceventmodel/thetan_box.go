package bceventmodel

type SignBuyBoxSignatureRequest struct {
	Sender       string `json:"sender"`
	PaymentERC20 string `json:"paymentERC20"`
	Price        string `json:"price"`
	BoxType      uint64 `json:"boxType"`
}

type SignBuyBoxSignatureResponse struct {
	Price        string `json:"price"`
	ExpiredAt    uint64 `json:"expiredAt"`
	Signature    string `json:"signature"`
	Sender       string `json:"sender"`
	BlockchainID uint64 `json:"blockchainId"`
	PaymentERC20 string `json:"paymentERC20"`
	BoxType      uint64 `json:"boxType"`
	BoxID        int64  `json:"boxId"`
}
