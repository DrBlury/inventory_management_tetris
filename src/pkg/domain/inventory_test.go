package domain

import (
	"testing"
)

func Test_inventoryItemPlaceCheck(t *testing.T) {
	exampleInventory := NewInventory(10, 10)
	exampleItem := &Item{
		ID:          1,
		Name:        "Test Item",
		Description: "This is a test item",
		Shape: Shape{
			Width:  4,
			Height: 3,
			Matrix: [][]int{
				{1, 0, 0, 0},
				{1, 1, 1, 1},
				{1, 0, 0, 0},
			},
		},
		SellValue: 100,
		BuyValue: 100,
	}

	exampleInventory.AddItemAtPosition(*exampleItem, &Position{X: 0, Y: 0, Rotation: 0})

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
	exampleInventory := NewInventory(10, 10)
	exampleItem := &Item{
		ID:          1,
		Name:        "Test Item",
		Description: "This is a test item",
		Shape: Shape{
			Width:  3,
			Height: 3,
			Matrix: [][]int{
				{1, 0, 0},
				{1, 1, 1},
				{1, 0, 0},
			},
		},
		SellValue: 100,
		BuyValue:  100,
	}

	exampleInventory.AddItemAtPosition(*exampleItem, &Position{X: 0, Y: 0, Rotation: 0})

	// add a few more items
	exampleInventory.AddItemAtPosition(*exampleItem, &Position{X: 4, Y: 7, Rotation: 0})
	// rotated item
	exampleItem.Shape.rotateCW()
	exampleInventory.AddItemAtPosition(*exampleItem, &Position{X: 7, Y: 7, Rotation: 0})

	exampleInventory.AddItemAtPosition(*createBox(6, 6, false), &Position{X: 4, Y: 0, Rotation: 0})

	// add another box shaped item
	exampleItem.Shape = Shape{
		Width:  2,
		Height: 2,
		Matrix: [][]int{
			{1, 1},
			{1, 1},
		},
	}
	error := exampleInventory.AddItemAtPosition(*exampleItem, &Position{X: 6, Y: 2, Rotation: 0})
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

	return &Item{
		ID:          1,
		Name:        "Box",
		Description: "A box",
		Shape: Shape{
			Width:  width,
			Height: height,
			Matrix: matrix,
		},
		SellValue: 100,
		BuyValue:  100,
	}
}

func createLShapedItem() *Item {
	return &Item{
		ID:          2,
		Name:        "L Shaped Item",
		Description: "An L shaped item",
		Shape: Shape{
			Width:  3,
			Height: 3,
			Matrix: [][]int{
				{1, 0, 0},
				{1, 0, 0},
				{1, 1, 1},
			},
		},
		SellValue: 100,
		BuyValue:  100,
	}
}

func Test_inventoryAddItem(t *testing.T) {
	exampleInventory := NewInventory(10, 10)
	if exampleInventory.AddItem(*createBox(3, 3, true)) != true {
		t.Errorf("expected true, got false")
		t.FailNow()
	}

	lShapedItem := createLShapedItem()
	// add 3 L shaped items
	for i := 0; i < 3; i++ {
		if exampleInventory.AddItem(*lShapedItem) != true {
			t.Errorf("expected true, got false")
			t.FailNow()
		}
	}

	// add a 6x6 box that's hollow
	if ok := exampleInventory.AddItem(*createBox(6, 6, false)); !ok {
		t.FailNow()
	}

	// add three 2x2 boxes
	for i := 0; i < 3; i++ {
		if ok := exampleInventory.AddItem(*createBox(2, 2, true)); !ok {
			t.FailNow()
		}
	}
}
