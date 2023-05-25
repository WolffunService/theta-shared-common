package rivalitemenum

type ItemType string

const (
	ITNone ItemType = ""

	// Minion
	ITSkin ItemType = "skin"

	// Cosmetic Profile
	ITAvatar      ItemType = "avatar"
	ITAvatarFrame ItemType = "avatar-frame"
	ITEmoticon    ItemType = "emoticon"
	ITNameColor   ItemType = "name-color"

	ITRivalBox ItemType = "rival-box"
	ITBigBox   ItemType = "big-box"
	ITEventBox ItemType = "event-box"
	ITMegaBox  ItemType = "mega-box"

	// Cosmetic Add-in
	ITGlow      ItemType = "glow"
	ITVoice     ItemType = "voice"
	ITDance     ItemType = "dance"
	ITFootprint ItemType = "footprint"
	ITCurtain   ItemType = "curtain"
	ITBackBling ItemType = "back-bling"
	ITFlyCraft  ItemType = "fly-craft"
	ITSpray     ItemType = "spray"
	ITVehicle   ItemType = "vehicle"

	// Currency
	ITThetanCoin ItemType = "thetan-coin"
	ITThetanGem  ItemType = "thetan-gem"

	ITEnhancer  ItemType = "enhancer"
	ITRivalBuck ItemType = "rival-buck"
	ITGold      ItemType = "gold"
	ITTrophy    ItemType = "trophy"

	ITSeasonPoint   ItemType = "season-point"
	ITSeasonBooster ItemType = "season-booster"
	ITEXP           ItemType = "exp"
)

func (itemType ItemType) IsCurrencyItem() bool {
	switch itemType {
	case ITThetanCoin, ITThetanGem,
		ITEnhancer, ITRivalBuck, ITGold, ITSeasonPoint, ITSeasonBooster:
		return true
	}
	return false
}

func (itemType ItemType) IsProfile() bool {
	switch itemType {
	case ITAvatar, ITAvatarFrame,
		ITEmoticon, ITNameColor:
		return true
	}
	return false
}

// hardcode truoc, move vao categories sau
func (itemType ItemType) IsAddIn() bool {
	switch itemType {
	case ITGlow,
		ITVoice,
		ITDance,
		ITFootprint,
		ITCurtain,
		ITBackBling,
		ITFlyCraft,
		ITSpray,
		ITVehicle:
		return true
	}
	return false
}
