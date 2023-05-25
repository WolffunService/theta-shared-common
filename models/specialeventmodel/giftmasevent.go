package specialeventmodel

func (GiftmasWheelEventModel) CollectionName() string {
	return "ServerSpecialEventData"
}

type GiftmasWheelEventModel struct {
	BaseEvent `json:",inline" bson:",inline"`
}

type GiftmasWheelEventModelClient struct {
	GiftmasWheelEventModel `json:",inline" bson:",inline"`
	EndIn                  int64 `json:"endIn" bson:"endIn"`
}

func (c GiftmasWheelEventModel) ToClient() *GiftmasWheelEventModelClient {
	clientData := &GiftmasWheelEventModelClient{
		GiftmasWheelEventModel: c,
		EndIn:                  c.EndIn(),
	}
	return clientData
}
