package domain

import (
	"fmt"
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
		Value: 100,
	}

	exampleInventory.AddItemAtPosition(*exampleItem, Position{X: 0, Y: 0})

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
			if exampleInventory.CheckItemPlacement(exampleItem, tc.position) != tc.expected {
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
		Value: 100,
	}

	exampleInventory.AddItemAtPosition(*exampleItem, Position{X: 0, Y: 0})

	// add a few more items
	exampleInventory.AddItemAtPosition(*exampleItem, Position{X: 4, Y: 7})
	// rotated item
	exampleItem.Shape.rotateCW()
	exampleInventory.AddItemAtPosition(*exampleItem, Position{X: 7, Y: 7})

	exampleInventory.AddItemAtPosition(*createBox(6, 6, false), Position{X: 4, Y: 0})

	// add another box shaped item
	exampleItem.Shape = Shape{
		Width:  2,
		Height: 2,
		Matrix: [][]int{
			{1, 1},
			{1, 1},
		},
	}
	exampleInventory.AddItemAtPosition(*exampleItem, Position{X: 6, Y: 2})

	if len(exampleInventory.Items) != 1 {
		t.Errorf("expected 1, got %d", len(exampleInventory.Items))
	}

	exampleInventory.PrintInventory()

	t.FailNow()
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
		Value: 100,
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
		Value: 100,
	}
}

func Test_inventoryAddItem(t *testing.T) {
	exampleInventory := NewInventory(10, 10)
	if exampleInventory.AddItem(*createBox(3, 3, true)) != true {
		t.Errorf("expected true, got false")
		t.FailNow()
	}

	exampleInventory.PrintInventory()

	lShapedItem := createLShapedItem()
	// add 3 L shaped items
	for i := 0; i < 3; i++ {
		if exampleInventory.AddItem(*lShapedItem) != true {
			t.Errorf("expected true, got false")
			t.FailNow()
		}
	}

	exampleInventory.PrintInventory()

	// add a big hollow box shaped item
	a6by6Box := createBox(6, 6, false)
	if exampleInventory.AddItem(*a6by6Box) != true {
		t.Errorf("expected true, got false")
		t.FailNow()
	}

	exampleInventory.AddItem(*createBox(3, 3, true))

	exampleInventory.AddItem(*createBox(2, 2, true))
	exampleInventory.AddItem(*createBox(2, 2, true))
	exampleInventory.PrintInventory()
	fmt.Println("")
	fmt.Println("---------------")
	fmt.Println("")

	exampleInventory.AddItem(*createBox(2, 2, true))

	exampleInventory.PrintInventory()
	fmt.Println("")
	fmt.Println("---------------")
	fmt.Println("")

	// add 4x2 box
	exampleInventory.AddItem(*createBox(4, 2, true))
	exampleInventory.PrintInventory()
	t.FailNow()

}
