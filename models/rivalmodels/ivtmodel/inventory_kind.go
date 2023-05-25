package ivtmodel

// InventoryKind kind of inventory, the biggest group
type InventoryKind uint8

const (
	IKCurrency InventoryKind = 1 + iota
	IKBox
	IKCosmeticProfile
	IKCosmeticAddIn
	IKMinion

	IKUndefined InventoryKind = 0
)
