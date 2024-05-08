package domain

import "fmt"

func NewInventory(inventoryMeta *InventoryMeta) *Inventory {
	return &Inventory{
		InventoryMeta: inventoryMeta,
		Items:         make([]*InventoryItem, 0),
	}
}

func (i *Inventory) AddItemAtPosition(item *Item, position *Position, quantity int64, durability int64) (*InventoryItem, error) {
	if !i.CheckItemPlacement(item, position) {
		return nil, &NoFitPositionError{}
	}

	inventoryItem := &InventoryItem{
		Item:           item,
		Position:       position,
		Quantity:       quantity,
		DurabilityLeft: durability,
	}

	i.Items = append(i.Items, inventoryItem)
	return inventoryItem, nil
}

func (i *Inventory) AddItem(item *Item, quantity int64, durability int64) bool {
	positionSuggestion, err := i.GetFitPosition(item)
	if err != nil {
		return false // item was not added because it did not fit
	}

	// add the item to the inventory
	i.Items = append(i.Items, &InventoryItem{
		Item:           item,
		Position:       positionSuggestion,
		Quantity:       quantity,
		DurabilityLeft: durability,
	})

	return true // item was added
}

func (i *Inventory) RemoveItem(itemId int64) {
	for idx, inventoryItem := range i.Items {
		if inventoryItem.Item.ItemMeta.Id == itemId {
			i.Items = append(i.Items[:idx], i.Items[idx+1:]...)
		}
	}
}

// GetFitPosition returns the first position where the item fits into the inventory
func (i *Inventory) GetFitPosition(item *Item) (*Position, error) {
	originalRotation := item.ItemMeta.Shape
	const POSSIBLE_ROTATIONS = 3 // 4 possible rotations (No rotation, 90, 180, 270 degrees)

	// for every possible cell, test all the possible rotations
	var maybePosition Position
	var tempShape *Shape
	for y := 0; y < int(i.InventoryMeta.Height); y++ {
		for x := 0; x < int(i.InventoryMeta.Width); x++ {
			tempShape = originalRotation
			// check item placement for every rotation
			for rotation := 0; rotation < POSSIBLE_ROTATIONS; rotation++ {
				if rotation > 0 {
					tempShape.applyRotations(1)
				}

				// Overwrite the maybePosition with the new values
				maybePosition.X = int64(x)
				maybePosition.Y = int64(y)
				maybePosition.Rotation = int64(rotation)

				if i.CheckItemPlacement(item, &maybePosition) {
					return &maybePosition, nil
				}
			}
		}
	}

	return nil, &NoFitPositionError{}
}

func (i *Inventory) getItemsInMatrix() [][]int {
	// create a temporary inventory matrix
	tempInventoryMatrix := make([][]int, i.InventoryMeta.Height)
	for column := range tempInventoryMatrix {
		tempInventoryMatrix[column] = make([]int, i.InventoryMeta.Width)
	}

	// place all items matrixes into the temporary inventory matrix
	for itemIdx, inventoryItem := range i.Items {
		// place the item into the temporary inventory matrix

		for y, row := range inventoryItem.Item.ItemMeta.Shape.getMatrix() {
			for x, cell := range row {
				// only add the item and not the empty cells
				if cell == 1 {
					tempInventoryMatrix[inventoryItem.Position.Y+int64(y)][inventoryItem.Position.X+int64(x)] = itemIdx + 1
				}
			}
		}
	}
	return tempInventoryMatrix
}

func (i *Inventory) CheckItemPlacement(item *Item, position *Position) bool {
	// Check if the item fits into the inventory or would reach out of bounds
	if position.X+item.ItemMeta.Shape.Width > i.InventoryMeta.Width || position.Y+item.ItemMeta.Shape.Height > i.InventoryMeta.Height {
		return false
	}

	tempInventoryMatrix := i.getItemsInMatrix()

	// temporarily apply the rotation to the item shape
	tempItemShape := item.ItemMeta.Shape
	tempItemShape.applyRotations(int(position.Rotation))

	// check if the item can be placed into the inventory
	for y, row := range tempItemShape.getMatrix() {
		for x, cell := range row {
			cellIsUsed := cell != 0
			inventoryCellIsUsed := tempInventoryMatrix[position.Y+int64(y)][position.X+int64(x)] != 0
			if cellIsUsed && inventoryCellIsUsed {
				return false
			}
		}
	}

	// TODO check for weight and max stack

	return true
}

func (i *Inventory) PrintInventory() {
	tempInventoryMatrix := i.getItemsInMatrix()
	fmt.Println("--------------------------------")
	fmt.Println("Inventory Items contained:")
	for idx, inventoryItem := range i.Items {
		fmt.Printf("%d: %s; Position: %d,%d\n", idx+1, inventoryItem.Item.Name, inventoryItem.Position.X, inventoryItem.Position.Y)
	}
	fmt.Println("Inventory:")

	// print the temporary inventory matrix
	prettyPrintMatrix(tempInventoryMatrix)
	fmt.Println(" ")

}
