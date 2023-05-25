package auditconst

// event name const
var EventName = struct {
	EN_ITEM_RECEIVED  string
	EN_TRADING_ITEM   string
	EN_BOX_OPEN       string
	EN_BOX_PURCHASE   string
	EN_COSMETIC_EVENT string
}{
	EN_ITEM_RECEIVED: "ItemReceived",
	EN_TRADING_ITEM:  "TradingItem",

	EN_BOX_OPEN: "OpenBoxItem",
	// EN_BOX_PURCHASE: "PurchaseBoxItem",

	EN_COSMETIC_EVENT: "CosmeticEvent",
}

// collection name const
var CollName = struct {
	CN_COSMETIC_AUDITS string
}{
	CN_COSMETIC_AUDITS: "CosmeticAudits",
}
