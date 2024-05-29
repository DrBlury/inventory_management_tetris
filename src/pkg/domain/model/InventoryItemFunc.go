package domain

import "fmt"

func (i *InventoryItem) CurrentShape() *Shape {
	shape := i.Item.ItemMeta.Shape
	shape.applyRotations(int(i.Position.Rotation))
	return shape
}

func (i *InventoryItem) RotateCW(amount int) {
	i.Position.Rotation = (i.Position.Rotation + int64(amount)) % 4
	fmt.Println("New rotation: ", i.Position.Rotation)
}
