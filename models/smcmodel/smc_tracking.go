package smcmodel

import (
	"github.com/WolffunService/thetan-shared-common/enums/smcenum"
	"github.com/func25/mongofunc/mocom"
)

type SmcMonitor struct {
	mocom.ID `bson:",inline"`
	Address  string                        `bson:"address"`
	Name     string                        `bson:"name"`
	Block    map[smcenum.ScanOffset]uint64 `bson:"block"`
}

func (SmcMonitor) CollName() string {
	return "SmcMonitors"
}
