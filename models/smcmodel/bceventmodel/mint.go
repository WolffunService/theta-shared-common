package bceventmodel

// ----- Requests & Responses -----

// sign mint
type MintSignReq struct {
	Sender string `json:"sender" binding:"required"`
	ItemID string `json:"itemId" binding:"required"`
}

type MintSignResp struct {
	Sender    string `json:"sender"`
	ExpiredAt int64  `json:"expiredAt"`
	TokenID   uint64 `json:"tokenId"`
	BID       uint64 `json:"bid"`
	Signature string `json:"signature"`
}

// sign multi mint
type MintMultiSignReq struct {
	Sender  string   `json:"sender" binding:"required"`
	ItemIDs []string `json:"itemIds" binding:"required"`
}

type MintMultiSignResp struct {
	Sender    string                  `json:"sender"`
	BID       uint64                  `json:"bid"`
	ExpiredAt int64                   `json:"expiredAt"`
	Signature string                  `json:"signatuer"`
	Tokens    []MintMultiSignRespUnit `json:"tokens"`
}

type MintMultiSignRespUnit struct {
	ItemID  string `json:"itemId"`
	TokenID uint64 `json:"tokenId"`
}

// admin mint
type MintAdminReq struct {
	ToAddress string `json:"toAddress" binding:"required"`
	ItemID    string `json:"itemId" binding:"required"`
}

type MintAdminResp struct {
	TokenID uint64 `json:"tokenId"`
}

// admin multi mint
type MintMultiAdminReq struct {
	NftInfos []MintAdminReq `json:"nftInfos" binding:"required"`
}

// ----- Events -----
type MintNFTEvent struct {
	TokenID         uint64
	ToAddress       string
	Block           uint64
	TransactionHash string
	BID             uint64
}

type MintMultiNFTsEvent struct {
	TokenIDs        []uint64
	ToAddress       string
	Block           uint64
	TransactionHash string
	BID             uint64
}

type MintMultiNFTsAdminEvent struct {
	Block           uint64
	TransactionHash string
	Data            []MintMultiNFTsAdminUnit
}

type MintMultiNFTsAdminUnit struct {
	To      string
	TokenID uint64
}

// -- timeout
type MintNFTTimeout struct {
	TokenIDs []uint64
}
