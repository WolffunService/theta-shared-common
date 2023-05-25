package specialeventmodel

func (SamsungEventModel) CollectionName() string {
	return "ServerSpecialEventData"
}

type SamsungEventModel struct {
	BaseEvent `json:",inline" bson:",inline"`
	WarmUpIn  int64 `json:"warmUpIn" bson:"warmUpIn"`
}

type SamsungEventModelClient struct {
	SamsungEventModel `json:",inline" bson:",inline"`
	EndIn             int64 `json:"endIn" bson:"endIn"`
}

func (c SamsungEventModel) ToClient() *SamsungEventModelClient {
	clientData := &SamsungEventModelClient{
		SamsungEventModel: c,
		EndIn:             c.EndIn(),
	}
	return clientData
}
