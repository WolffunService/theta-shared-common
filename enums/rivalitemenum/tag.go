package rivalitemenum

type Tag string

const (
	ITagAddInNone Tag = "none"

	ITagCurrencyCrypto Tag = "crypto-currency"
	ITagCurrencyLocal  Tag = "local-currency"

	ITagCosmeticTheatrics  Tag = "theatrics"
	ITagCosmeticEquipments Tag = "equipments"

	ITagProfileAvatar      Tag = "avatar"
	ITagProfileAvatarFrame Tag = "avatar-frame"
	ITagProfileEmoticon    Tag = "emoticon"
	ITagProfileNameColor   Tag = "name-color"

	ITagAddInCurtain   Tag = "curtain"
	ITagAddInGlow      Tag = "glow"
	ITagAddInVoice     Tag = "voice"
	ITagAddInDance     Tag = "dance"
	ITagAddInFootprint Tag = "footprint"
	ITagAddInSpray     Tag = "spray"
	ITagAddInBackBling Tag = "back-bling"
	ITagAddInFlyCraft  Tag = "fly-craft"
	ITagAddInVehicle   Tag = "vehicle"
)
