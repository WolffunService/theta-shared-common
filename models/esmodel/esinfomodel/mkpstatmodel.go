package esinfomodel

import (
	"time"
)

type MKPStatModel struct {
	Action       int       `json:"action"` // Sale, Sold, Cancel, Buy BOX, Fusion, Buy Private sale, WithDraw, Upgrade, Deposit, Buy ticket special event
	RefType      int       `json:"ref_type"`
	RefId        string    `json:"ref_id"`
	OnMarketTime time.Time `json:"on_market_time"`
	THCPrice     int64     `json:"thc_price"`
	USDPrice     int64     `json:"usd_price"`
	THCFee       int64     `json:"thc_fee"`
	THGFee       int64     `json:"thg_fee"`
}
