package domain

import "fmt"

type Inventory struct {
	InventoryMeta InventoryMeta
	Items         []InventoryItem
}

func NewInventory(inventoryMeta InventoryMeta) *Inventory {
	return &Inventory{
		InventoryMeta: inventoryMeta,
		Items:         make([]InventoryItem, 0),
	}
}

func (i *Inventory) AddItemAtPosition(item Item, position *Position, quantity int, durability int) (*InventoryItem, error) {
	if !i.CheckItemPlacement(&item, position) {
		return nil, &NoFitPositionError{}
	}

	inventoryItem := &InventoryItem{
		Item:           item,
		Position:       *position,
		Quantity:       quantity,
		DurabilityLeft: durability,
	}

	i.Items = append(i.Items, *inventoryItem)
	return inventoryItem, nil
}

func (i *Inventory) AddItem(item Item, quantity int, durability int) bool {
	positionSuggestion, err := i.GetFitPosition(item)
	if err != nil {
		return false // item was not added because it did not fit
	}

	// rotate the item if necessary
	for rotation := 0; rotation < positionSuggestion.Rotation; rotation++ {
		item.ItemMeta.Shape = rotateCW(item.ItemMeta.Shape)
	}

	// add the item to the inventory
	i.Items = append(i.Items, InventoryItem{
		Item:           item,
		Position:       *positionSuggestion,
		Quantity:       quantity,
		DurabilityLeft: durability,
	})

	return true // item was added
}

func (i *Inventory) RemoveItem(itemID int) {
	for idx, inventoryItem := range i.Items {
		if inventoryItem.Item.ItemMeta.ID == itemID {
			i.Items = append(i.Items[:idx], i.Items[idx+1:]...)
		}
	}
}

// GetFitPosition returns the first position where the item fits into the inventory
func (i *Inventory) GetFitPosition(item Item) (*Position, error) {
	originalRotation := item.ItemMeta.Shape
	const POSSIBLE_ROTATIONS = 3 // 4 possible rotations (No rotation, 90, 180, 270 degrees)

	// for every possible cell, test all the possible rotations
	var maybePosition Position
	tempShape := item.ItemMeta.Shape
	for y := 0; y < i.InventoryMeta.Height; y++ {
		for x := 0; x < i.InventoryMeta.Width; x++ {
			tempShape = originalRotation
			// check item placement for every rotation
			for rotation := 0; rotation < POSSIBLE_ROTATIONS; rotation++ {
				if rotation > 0 {
					// TODO Maybe there's a bug here?
					tempShape = rotateCW(tempShape)
				}

				// Overwrite the maybePosition with the new values
				maybePosition.X = x
				maybePosition.Y = y
				maybePosition.Rotation = rotation

				if i.CheckItemPlacement(&item, &maybePosition) {
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

		for y, row := range inventoryItem.Item.ItemMeta.Shape.Matrix {
			for x, cell := range row {
				// only add the item and not the empty cells
				if cell == 1 {
					tempInventoryMatrix[inventoryItem.Position.Y+y][inventoryItem.Position.X+x] = itemIdx + 1
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
	var rotatedShape Shape
	// temporarily apply the rotation to the item shape
	for rotation := 0; rotation < position.Rotation; rotation++ {
		rotatedShape = rotateCW(item.ItemMeta.Shape)
	}

	// check if the item can be placed into the inventory
	for y, row := range rotatedShape.Matrix {
		for x, cell := range row {
			cellIsUsed := cell != 0
			inventoryCellIsUsed := tempInventoryMatrix[position.Y+y][position.X+x] != 0
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
