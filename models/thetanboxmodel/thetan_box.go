package thetanboxmodel

import "github.com/WolffunService/theta-shared-common/enums/thetanboxenum"

type BoxData struct {
	BoxType thetanboxenum.BoxType `json:"boxType" bson:"boxType"`
	Amount  int                   `json:"amount" bson:"amount"`
}
