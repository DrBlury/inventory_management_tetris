package domain

import (
	"testing"
)

func genExampleInventoryMeta() *InventoryMeta {
	exampleInventoryMeta := &InventoryMeta{
		Id:        1,
		Name:      "Test Inventory",
		UserId:    1,
		Width:     10,
		Height:    10,
		MaxWeight: 100,
	}
	return exampleInventoryMeta
}

func genExampleItemWithMatrix(matrix [][]int) *Item {
	// create raw shape using the matrix
	rawShape := ""
	for _, row := range matrix {
		for _, cell := range row {
			if cell == 1 {
				rawShape += "#"
			} else {
				rawShape += "."
			}
		}
	}
	return &Item{
		ItemMeta: &ItemMeta{
			Id: 1,
			Shape: &Shape{
				Width:    int64(len(matrix[0])),
				Height:   int64(len(matrix)),
				RawShape: rawShape,
			},
			Weight:   10,
			MaxStack: 1,
		},

		Name:       "Test Item",
		Text:       "This is a test item",
		SellValue:  100,
		BuyValue:   100,
		Durability: 100,
		Variant:    "Test Variant",
		Type:       ItemType_ARMOR,
	}
}

func Test_inventoryItemPlaceCheck(t *testing.T) {

	exampleInventory := NewInventory(genExampleInventoryMeta())
	exampleItem := genExampleItemWithMatrix([][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
	})

	_, err := exampleInventory.AddItemAtPosition(exampleItem, &Position{X: 0, Y: 0, Rotation: 0}, 1, 100)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
		t.FailNow()
	}

	testCases := []struct {
		name     string
		position Position
		expected bool
	}{
		{
			name:     "placeable",
			position: Position{X: 4, Y: 4},
			expected: true,
		},
		{
			name:     "out of bounds",
			position: Position{X: 9, Y: 2},
			expected: false,
		},
		{
			name:     "overlapping",
			position: Position{X: 1, Y: 1},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if exampleInventory.CheckItemPlacement(exampleItem, &tc.position) != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, !tc.expected)
			}
		})
	}
}

func Test_inventoryItemAdd(t *testing.T) {
	exampleInventory := NewInventory(genExampleInventoryMeta())
	exampleItem := genExampleItemWithMatrix([][]int{
		{1, 0, 0},
		{1, 1, 1},
		{1, 0, 0},
	})

	_, _ = exampleInventory.AddItemAtPosition(exampleItem, &Position{X: 0, Y: 0, Rotation: 0}, 1, 100)

	// add a few more items
	_, _ = exampleInventory.AddItemAtPosition(exampleItem, &Position{X: 4, Y: 7, Rotation: 0}, 1, 100)

	// rotated item
	_, _ = exampleInventory.AddItemAtPosition(exampleItem, &Position{X: 7, Y: 7, Rotation: 1}, 1, 100)

	_, _ = exampleInventory.AddItemAtPosition(createBox(6, 6, false), &Position{X: 4, Y: 0, Rotation: 0}, 1, 100)

	// add another box shaped item
	exampleItem.ItemMeta.Shape = &Shape{
		Width:    2,
		Height:   2,
		RawShape: "####",
	}
	_, error := exampleInventory.AddItemAtPosition(exampleItem, &Position{X: 6, Y: 2, Rotation: 0}, 1, 100)
	if error != nil {
		t.Errorf("expected nil, got %v", error)
		t.FailNow()
	}

	if len(exampleInventory.Items) != 5 {
		t.Errorf("expected 1, got %d", len(exampleInventory.Items))
		t.FailNow()
	}
}

func createBox(width int, height int, filled bool) *Item {

	// Create empty matrix
	matrix := make([][]int, height)
	for i := range matrix {
		matrix[i] = make([]int, width)
	}

	// Fill the matrix if the box is filled
	if filled {
		for i := range matrix {
			for j := range matrix[i] {
				matrix[i][j] = 1
			}
		}
	}

	// Else draw a hollow box
	if !filled {
		for i := range matrix {
			for j := range matrix[i] {
				if i == 0 || i == height-1 || j == 0 || j == width-1 {
					matrix[i][j] = 1
				}
			}
		}
	}

	exampleItem := genExampleItemWithMatrix(matrix)
	exampleItem.Name = "Box"
	return exampleItem
}

func createLShapedItem() *Item {
	return genExampleItemWithMatrix([][]int{
		{1, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
	})
}

func Test_inventoryAddItem(t *testing.T) {
	exampleInventory := NewInventory(genExampleInventoryMeta())
	_, ok := exampleInventory.AddItem(createBox(3, 3, true), 1, 100)
	if !ok {
		t.Errorf("expected true, got false")
		t.FailNow()
	}

	lShapedItem := createLShapedItem()
	// add 3 L shaped items
	for i := 0; i < 3; i++ {
		_, ok := exampleInventory.AddItem(lShapedItem, 1, 100)
		if !ok {
			t.Errorf("expected true, got false")
			t.FailNow()
		}
	}

	// add a 6x6 box that's hollow
	_, ok = exampleInventory.AddItem(createBox(6, 6, false), 1, 100)
	if !ok {
		t.FailNow()
	}

	// add three 2x2 boxes
	for i := 0; i < 3; i++ {
		_, ok := exampleInventory.AddItem(createBox(2, 2, true), 1, 100)
		if !ok {
			t.FailNow()
		}
	}
}
