package cosmeticcreatingmodel

import "github.com/WolffunService/thetan-shared-common/enums/cosmeticenum"

type RivalsCreateCosmeticRequest struct {
	TypeId int                              `json:"typeId" bson:"typeId"`
	UserId string                           `json:"userId" bson:"userId"`
	Status cosmeticenum.CosmeticStatus      `json:"status" bson:"status"`
	Source cosmeticenum.CosmeticEventSource `json:"source" bson:"source"`
}
