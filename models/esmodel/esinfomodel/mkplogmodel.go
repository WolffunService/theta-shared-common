package esinfomodel

type MkpLogModel struct {
	UserId          string `json:"user_id"`
	UserAddress     string `json:"user_address"`
	UserName        string `json:"user_name"`
	UserCountry     string `json:"user_country"`
	Action          int    `json:"action"` // Sell, Buy, Cancel, Buy BOX, Fusion, Buy Private sale, WithDraw, Upgrade, Deposit, Buy ticket special event
	RefType         int    `json:"ref_type"`
	RefId           string `json:"ref_id"`
	THCPrice        int64  `json:"thc_price"`
	USDPrice        int64  `json:"usd_price"`
	PartnerId       string `json:"partner_id"`
	PartnerAddress  string `json:"partner_address"`
	PartnerUserName string `json:"partner_user_name"`
	PartnerCountry  string `json:"partner_country"`
}
