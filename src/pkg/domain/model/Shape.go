package domain

import (
	"fmt"
)

type Shape struct {
	RawShape string
	Width    int
	Height   int
}

func (s *Shape) getMatrix() [][]int {
	// Initialize the shape matrix
	matrix := make([][]int, s.Height)
	for i := range matrix {
		matrix[i] = make([]int, s.Width)
	}

	// cut the raw shape into rows
	rows := []string{}
	for i := 0; i < len(s.RawShape); i += s.Width {
		rows = append(rows, s.RawShape[i:i+s.Width])
	}

	// fill the shape matrix
	for i, row := range rows {
		for j, cell := range row {
			if cell == '#' {
				matrix[i][j] = 1
			}
		}
	}

	return matrix
}

func (s *Shape) printShape() {
	for _, row := range s.getMatrix() {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("#") // Occupied spaces are represented by "#"
			} else {
				fmt.Print(".") // Empty spaces are represented by "."
			}
		}
		fmt.Println() // New line after each row
	}
}

func (s *Shape) applyRotations(rotations int) {
	rotatedMatrix := s.getRotatedMatrixCW(rotations)
	height := len(rotatedMatrix)
	width := len(rotatedMatrix[0])
	// generate new raw shape
	rawShape := ""
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if rotatedMatrix[i][j] == 1 {
				rawShape += "#"
			} else {
				rawShape += "."
			}
		}
	}
	s.Height = height
	s.Width = width
	s.RawShape = rawShape
}

// TODO maybe add amount of rotations to apply
func (s *Shape) getRotatedMatrixCW(rotations int) [][]int {
	// get matrix of shape
	itemMatrix := s.getMatrix()
	for i := 0; i < rotations; i++ {
		itemMatrix = rotateMatrixCW(itemMatrix)
	}
	return itemMatrix
}

func rotateMatrixCW(oldMatrix [][]int) [][]int {
	// Create a new matrix with the same dimensions as the original shape but with the width and height swapped
	width := len(oldMatrix[0])
	height := len(oldMatrix)

	newMatrix := make([][]int, width)
	for i := range newMatrix {
		newMatrix[i] = make([]int, height)
	}

	// Fill the new matrix with the rotated values (clockwise)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			newMatrix[j][height-i-1] = oldMatrix[i][j]
		}
	}

	return newMatrix
}
