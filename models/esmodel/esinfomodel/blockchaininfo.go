package esinfomodel

type BlockchainInfoModel struct {
	Block           int64  `json:"block"`
	TokenID         int64  `json:"token_id"`
	TransactionHash string `json:"transaction_hash"`
	PaymentToken    string `json:"payment_token"`
	BlockchainID    int64  `json:"blockchain_id"`
}
