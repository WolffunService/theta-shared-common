package bceventmodel

// ----- Requests & Responses ------
type SignSellingReq struct {
	NFTAddress          string `json:"nftAddress" binding:"required"`
	PaymentTokenAddress string `json:"paymentTokenAddress" binding:"required"`
	TokenID             int64  `json:"tokenId" binding:"required"`
	Price               string `json:"price" binding:"required"`
}

type SignSellingResp struct {
	Price               string `json:"price"`
	ExpiredAt           int64  `json:"expiredAt"`
	Signature           string `json:"signature"`
	NFTAddress          string `json:"nftAddress"`
	TokenID             int64  `json:"tokenId"`
	PaymentTokenAddress string `json:"paymentTokenAddress"`
	BID                 uint64 `json:"bid"`
}

// ----- Events -----

type BuyNFTEvent struct {
	Block           uint64
	TransactionHash string
	TokenID         uint64
	Price           string
	PaymentToken    string
	Seller          string
	Buyer           string
	Fee             uint64
}

type SellNFTEvent struct {
	Block           uint64
	TransactionHash string
	TokenID         uint64
	Price           string
	PaymentToken    string
	BID             uint64
}

type CancelSellNFTEvent struct {
	Block           uint64
	TransactionHash string
	TokenID         uint64
}

type TimeoutEvent struct {
	BID uint64
}
