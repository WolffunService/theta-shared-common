package cosmeticenum

type CosmeticType int // #enum: "Cosmetic Type"

const (
	CT_none CosmeticType = iota
	CT_AVATAR
	CT_AVATAR_FRAME
	CT_EMOTE
	CT_SPACE_SHIP

	CT_end
)

type CosmeticRarity int // #enum: "Cosmetic Rarity"

const (
	CR_COMMON    CosmeticRarity = 1
	CR_EPIC      CosmeticRarity = 2
	CR_LEGENDARY CosmeticRarity = 3
)

type CosmeticStatus int // #enum: "Cosmetic Status"

const (
	CS_FREE      CosmeticStatus = 1
	CS_MINTING   CosmeticStatus = 2
	CS_AVAILABLE CosmeticStatus = 3
	CS_SELLING   CosmeticStatus = 20
)

type CosmeticEventSource int // #enum: "Cosmetic Event Source"

const (
	CES_OPEN_BOX       CosmeticEventSource = 1
	CES_SEASON_END     CosmeticEventSource = 2
	CES_RANKING_REWARD CosmeticEventSource = 3
	CES_SOLD           CosmeticEventSource = 4
	CES_PURCHASED      CosmeticEventSource = 5
	CES_MINT           CosmeticEventSource = 6
	CES_ADMIN          CosmeticEventSource = 7
)
