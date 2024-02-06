package domain

import "fmt"

type InventoryItem struct {
	Item     Item
	Position Position
}

type Position struct {
	X int
	Y int
}

type PositionSuggestion struct {
	Position Position
	Rotation int
}

type Inventory struct {
	ID      int
	Name    string
	Owner   string
	OwnerID int
	Width   int
	Height  int
	Items   []*InventoryItem
}

func NewInventory(width int, height int) *Inventory {
	return &Inventory{
		Width:  width,
		Height: height,
		Items:  make([]*InventoryItem, 0),
	}
}

func (i *Inventory) AddItemAtPosition(item Item, position Position) {
	if i.CheckItemPlacement(&item, position) {
		i.Items = append(i.Items, &InventoryItem{
			Item:     item,
			Position: position,
		})
	}
}

func (i *Inventory) AddItem(item Item) bool {
	positionSuggestion, err := i.GetFitPosition(item)
	if err != nil {
		return false // item was not added because it did not fit
	}

	// rotate the item if necessary
	for rotation := 0; rotation < positionSuggestion.Rotation; rotation++ {
		item.Shape.rotateCW()
	}

	// add the item to the inventory
	i.Items = append(i.Items, &InventoryItem{
		Item:     item,
		Position: positionSuggestion.Position,
	})

	return true // item was added
}

func (i *Inventory) RemoveItem(item Item) {
	for idx, inventoryItem := range i.Items {
		if inventoryItem.Item.ID == item.ID {
			i.Items = append(i.Items[:idx], i.Items[idx+1:]...)
		}
	}
}

// GetFitPosition returns the first position where the item fits into the inventory
func (i *Inventory) GetFitPosition(item Item) (*PositionSuggestion, error) {
	originalRotation := item.Shape
	const POSSIBLE_ROTATIONS = 3 // 4 possible rotations (No rotation, 90, 180, 270 degrees)

	// for every possible cell, test all the possible rotations
	for y := 0; y < i.Height; y++ {
		for x := 0; x < i.Width; x++ {
			item.Shape = originalRotation
			// check item placement for every rotation
			for rotation := 0; rotation < POSSIBLE_ROTATIONS; rotation++ {
				if rotation > 0 {
					item.Shape.rotateCW()
				}

				if i.CheckItemPlacement(&item, Position{X: x, Y: y}) {
					return &PositionSuggestion{
						Position: Position{X: x, Y: y},
						Rotation: rotation,
					}, nil
				}
			}
		}
	}

	return nil, &NoFitPositionError{}
}

func (i *Inventory) getItemsInMatrix() [][]int {
	// create a temporary inventory matrix
	tempInventoryMatrix := make([][]int, i.Height)
	for column := range tempInventoryMatrix {
		tempInventoryMatrix[column] = make([]int, i.Width)
	}

	// place all items matrixes into the temporary inventory matrix
	for itemIdx, inventoryItem := range i.Items {
		// place the item into the temporary inventory matrix
		for y, row := range inventoryItem.Item.Shape.Matrix {
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

func (i *Inventory) CheckItemPlacement(item *Item, position Position) bool {
	// Check if the item fits into the inventory or would reach out of bounds
	if position.X+item.Shape.Width > i.Width || position.Y+item.Shape.Height > i.Height {
		return false
	}

	tempInventoryMatrix := i.getItemsInMatrix()

	// check if the item can be placed into the inventory
	for y, row := range item.Shape.Matrix {
		for x, cell := range row {
			cellIsUsed := cell != 0
			inventoryCellIsUsed := tempInventoryMatrix[position.Y+y][position.X+x] != 0
			if cellIsUsed && inventoryCellIsUsed {
				return false
			}
		}
	}

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

// prettyPrintMatrix prints a matrix in a human readable format
// It draws cells with outlines and fills them with their content
// Example:
// ---------
// | 1 | 2 |
// | 3 | 4 |
// ---------
// If the content is bigger than 1 character, all cells will be the size of the biggest content
func prettyPrintMatrix(matrix [][]int) {
	// find the biggest content
	biggestContent := 0
	for _, row := range matrix {
		for _, cell := range row {
			if cell > biggestContent {
				biggestContent = cell
			}
		}
	}

	// print the matrix
	fmt.Println("-----------------------")
	for _, row := range matrix {
		for _, cell := range row {
			fmt.Printf("| %*d ", len(fmt.Sprint(biggestContent)), cell)
		}
		fmt.Println("|")
	}
	fmt.Println("-----------------------")
}
