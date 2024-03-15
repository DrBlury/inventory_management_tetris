package domain

type InventoryItem struct {
	Item           Item
	Position       Position
	Quantity       int
	DurabilityLeft int
}
