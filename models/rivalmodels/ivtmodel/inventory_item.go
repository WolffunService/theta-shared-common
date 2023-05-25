package ivtmodel

// ItemData global item data response to client
type ItemData struct {
	Kind InventoryKind `json:"kind"`
	Data any           `json:"data"`
}

// Shortcut item data types
type (
	ListItemData []ItemData
	MapItemData  map[string]ItemData
)
