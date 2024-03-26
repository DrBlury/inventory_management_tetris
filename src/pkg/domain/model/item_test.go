package domain

import (
	"reflect"
	"testing"
)

// Test_parseShape tests the parseShape function.
func Test_parseShape(t *testing.T) {
	rawShape := "#####..##..#####"
	width := 4
	height := 4
	shape := &Shape{}
	shape.parseShape(rawShape, width, height)

	expectedShapeMatrix := [][]int{
		{1, 1, 1, 1},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{1, 1, 1, 1},
	}

	for i, row := range shape.Matrix {
		for j, cell := range row {
			if cell != expectedShapeMatrix[i][j] {
				t.Errorf("expected %d, got %d", expectedShapeMatrix[i][j], cell)
			}
		}
	}
}

// Test_printShape tests the printShape function.
func Test_printShape(t *testing.T) {
	shape := Shape{
		Matrix: [][]int{
			{1, 1, 1, 1},
			{1, 0, 0, 1},
			{1, 0, 0, 1},
			{1, 1, 1, 1},
		},
		Width:  4,
		Height: 4,
	}

	shape.printShape()
	t.FailNow()
}

// Test_rotateShape tests the currentShape function with rotation.
func Test_rotateShape(t *testing.T) {
	exampleItem := genExampleItemWithMatrix([][]int{
		{1, 1, 1},
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	})

	// create a real instance of the item
	inventoryItem := InventoryItem{
		Item:           *exampleItem,
		Position:       Position{X: 0, Y: 0, Rotation: 0},
		Quantity:       1,
		DurabilityLeft: 100,
	}

	// rotate the item
	inventoryItem.RotateCW(1)

	// get the rotated shape
	rotatedShape := inventoryItem.CurrentShape()

	expectedShape := [][]int{
		{0, 0, 0, 1},
		{1, 1, 1, 1},
		{0, 0, 0, 1},
	}

	if !reflect.DeepEqual(rotatedShape.Matrix, expectedShape) {
		t.Errorf("Rotation failed.\nExpected:\n%v\nGot:\n%v", expectedShape, rotatedShape.Matrix)
	}
}
