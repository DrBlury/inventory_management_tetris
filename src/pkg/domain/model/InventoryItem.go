package domain

import "fmt"

type InventoryItem struct {
	Item           Item
	Position       Position
	Quantity       int
	DurabilityLeft int
}

func (i *InventoryItem) CurrentShape() Shape {
	shape := i.Item.ItemMeta.Shape
	shape.applyRotations(i.Position.Rotation)
	return shape
}

func (i *InventoryItem) RotateCW(amount int) {
	i.Position.Rotation = (i.Position.Rotation + amount) % 4
	fmt.Println("New rotation: ", i.Position.Rotation)
}
