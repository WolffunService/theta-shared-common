package esmodel

import (
	"time"
)

//var UserIndicesMapping = esapi.IndicesPutMappingRequest{
//	Index: []string{"user"},
//	Body:  nil,
//}

//Nhóm stat battle, earning được track sau mỗi battle, loại stat này có dạng append-only time series nên sẽ sử dụng
//Data stream để index

//

type StatName string

// UserModel Các stat cố định của player sẽ define tại struct này để phục vụ filtering
type UserModel struct {
	ID              string    `json:"id"`
	Username        string    `json:"username"`
	Mail            string    `json:"mail"`
	Address         string    `json:"address"`
	Country         string    `json:"country"`
	Created         time.Time `json:"created"`
	GameOpened      time.Time `json:"game_opened"`
	WalletConnected time.Time `json:"wallet_connected"`
}

type PlayerMapping struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Mail     string    `json:"mail"`
	Country  string    `json:"country"`
	Created  time.Time `json:"created"`
}

func (PlayerMapping) Index() string {
	return "players"
}
