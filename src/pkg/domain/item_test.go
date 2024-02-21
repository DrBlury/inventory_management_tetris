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

// Test_rotateShapeCCW tests the rotateShape function.
func Test_rotateShapeCCW(t *testing.T) {
	tShapeMatrix := [][]int{
		{1, 1, 1},
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}

	// create shape
	tShape := Shape{Width: 3, Height: 4, Matrix: tShapeMatrix}

	tShape.rotateCCW()

	expectedShape := [][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
	}

	if !reflect.DeepEqual(tShape.Matrix, expectedShape) {
		t.Errorf("Rotation failed.\nExpected:\n%v\nGot:\n%v", expectedShape, tShape.Matrix)
	}
}

// Test_rotateShapeCW tests the rotateShape function.
func Test_rotateShapeCW(t *testing.T) {
	tShapeMatrix := [][]int{
		{1, 1, 1},
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}

	// create shape
	tShape := Shape{Width: 3, Height: 4, Matrix: tShapeMatrix}

	tShape.rotateCW()

	expectedShape := [][]int{
		{0, 0, 0, 1},
		{1, 1, 1, 1},
		{0, 0, 0, 1},
	}

	if !reflect.DeepEqual(tShape.Matrix, expectedShape) {
		t.Errorf("Rotation failed.\nExpected:\n%v\nGot:\n%v", expectedShape, tShape.Matrix)
	}
}

// Test_rotateShapeCWAndCCW tests the rotateShape function.
func Test_rotateShapeCWAndCCW(t *testing.T) {
	tShapeMatrix := [][]int{
		{1, 1, 1},
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}

	// create shape
	tShape := Shape{Width: 3, Height: 4, Matrix: tShapeMatrix}

	tShape.rotateCW()
	tShape.rotateCCW()

	if !reflect.DeepEqual(tShapeMatrix, tShape.Matrix) {
		t.Errorf("Rotation failed.\nExpected:\n%v\nGot:\n%v", tShape, tShape)
	}
}

// Test_doubleRotateOrFlip tests the doubleRotateOrFlip function.
func Test_doubleRotateOrFlip(t *testing.T) {
	tShape := [][]int{
		{1, 1, 1},
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}

	// create item
	tShapeItem := Shape{Width: 3, Height: 4, Matrix: tShape}

	tShapeItem.flip()

	expectedShape := [][]int{
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	if !reflect.DeepEqual(tShapeItem.Matrix, expectedShape) {
		t.Errorf("Rotation failed.\nExpected:\n%v\nGot:\n%v", expectedShape, tShapeItem.Matrix)
	}
}
