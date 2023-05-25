package internalmessage

type InternalMessage struct {
	ItemId    string `json:"itemId"`
	DataType  int    `json:"dataType"`
	Block     int    `json:"block"`
	MkpAction int    `json:"mkpAction"`
}
