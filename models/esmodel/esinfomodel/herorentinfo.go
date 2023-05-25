package esinfomodel

import "time"

type RentInfoModel struct {
	MaxRentalPeriod int       `json:"max_rental_period"`
	RentBattles     int       `json:"rent_battles"`
	ExpiredTime     time.Time `json:"expired_time"`
}
