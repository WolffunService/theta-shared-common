package specialeventmodel

func (CoinBaseEventModel) CollectionName() string {
	return "ServerSpecialEventData"
}

type CoinBaseEventModel struct {
	BaseEvent       `json:",inline" bson:",inline"`
	BannerURL       string `json:"bannerURL" bson:"bannerURL"`
	BannerMobileURL string `json:"bannerMobileURL" bson:"bannerMobileURL"`
	BannerDetailURL string `json:"bannerDetailURL" bson:"bannerDetailURL"`
	DirectURLEvent  string `json:"directURLEvent" bson:"directURLEvent"`
}

type CoinBaseEventModelClient struct {
	CoinBaseEventModel `json:",inline" bson:",inline"`
	EndIn              int64 `json:"endIn" bson:"endIn"`
}

func (c CoinBaseEventModel) ToClient() *CoinBaseEventModelClient {
	clientData := &CoinBaseEventModelClient{
		CoinBaseEventModel: c,
		EndIn:              c.EndIn(),
	}
	return clientData
}
