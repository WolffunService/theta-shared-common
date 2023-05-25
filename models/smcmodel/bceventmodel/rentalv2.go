package bceventmodel

// ----- Requests & Responses ------
type SignLendingReq struct {
	TokenID uint64 `json:"tokenId" binding:"required"`
	Price   string `json:"price" binding:"required"`
}

type SignLendingResp struct {
	Price     string `json:"price"`
	Signature string `json:"signature"`
	TokenID   uint64 `json:"tokenId"`
	BID       uint64 `json:"bid"`
	ExpiredAt int64  `json:"expiredAt"`
}

type SignUnlockReq struct {
	TokenID uint64 `json:"tokenId" binding:"required"`
	Version int    `json:"version"`
}

type SignUnlockResp struct {
	TokenID   uint64 `json:"tokenId"`
	BID       uint64 `json:"bid"`
	Old       bool   `json:"old"`
	Signature string `json:"signature"`
	ExpiredAt int64  `json:"expiredAt"`
}

// ----- Events -----
type LendingEvent struct {
	Block           uint64
	TransactionHash string
	TokenID         uint64
	Price           string
	BID             uint64
}

type RentingEvent struct {
	Block           uint64
	TransactionHash string
	TokenID         uint64
	Price           string
	Renter          string
	Fee             uint64
}

type CancelLendEvent struct {
	TokenID         uint64
	Block           uint64
	TransactionHash string
}

type UnlockHeroEvent struct {
	TokenID         uint64
	Block           uint64
	TransactionHash string
}
