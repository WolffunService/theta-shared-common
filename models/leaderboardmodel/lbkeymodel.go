package leaderboardmodel

import (
	"fmt"

	"github.com/WolffunService/thetan-shared-common/enums/leaderboardenum"
)

type LBKeyModel struct {
	key leaderboardenum.LBKeyPrefix
	id  int
}

func NewLBKeyModel(key leaderboardenum.LBKeyPrefix, id int) *LBKeyModel {
	return &LBKeyModel{key: key, id: id}
}

func (l *LBKeyModel) String() string {
	return fmt.Sprintf("%s_%d", l.key, l.id)
}

func (l *LBKeyModel) GetID() int {
	return l.id
}

func (l *LBKeyModel) GetKey() string {
	return string(l.key)
}
