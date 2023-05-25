package thetanrivalminionenum

import (
	"github.com/WolffunService/thetan-shared-common/enums/rivalitemenum"
	"github.com/WolffunService/thetan-shared-common/enums/thetanrivalerrorenum/thetanrivalbotenum"
	"github.com/WolffunService/thetan-shared-common/enums/thetanrivalerrorenum/thetanrivalconst"
)

type MinionBot struct {
	Type   int                           `json:"type"`
	Name   string                        `json:"name"`
	Rarity thetanrivalconst.MinionRarity `json:"rarity"`
	Skins  []MinionSkinBot               `json:"skins"`
}

type MinionSkinBot struct {
	Id          int                               `json:"id"`
	Name        string                            `json:"name"`
	Rarity      thetanrivalconst.MinionSkinRarity `json:"rarity"`
	UniformId   int                               `json:"uniformId"`
	ListBotRate []thetanrivalbotenum.BotRate      `json:"listBotRate"`
}

type CosmeticBot struct {
	Name         string                          `json:"name,omitempty"`
	SkinId       int                             `json:"skinId"`
	CosmeticType rivalitemenum.Tag               `json:"cosmeticType"`
	Rarity       thetanrivalconst.CosmeticRarity `json:"rarity"`
	ListBotRate  []thetanrivalbotenum.BotRate    `json:"listBotRate"`
}
