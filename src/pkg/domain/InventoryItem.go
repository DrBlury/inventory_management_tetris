package domain

type InventoryItem struct {
	Item           Item
	Position       Position
	Quantity       int
	DurabilityLeft int
}

func (i *InventoryItem) CurrentShape() Shape {
	shape := i.Item.ItemMeta.Shape
	for rotation := 0; rotation < i.Position.Rotation; rotation++ {
		shape = rotateCW(shape)
	}
	return shape
}

func (i *InventoryItem) RotateCW(amount int) {
	i.Position.Rotation = (i.Position.Rotation + amount) % 4
}
