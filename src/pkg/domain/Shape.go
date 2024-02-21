package domain

import "fmt"

type Shape struct {
	RawShape string
	Width  int
	Height int
	Matrix [][]int
}

func (s *Shape) parseShape(rawShape string, width int, height int) {
	// Initialize the shape matrix
	matrix := make([][]int, height)
	for i := range matrix {
		matrix[i] = make([]int, width)
	}

	// cut the raw shape into rows
	rows := []string{}
	for i := 0; i < len(rawShape); i += width {
		rows = append(rows, rawShape[i:i+width])
	}

	// fill the shape matrix
	for i, row := range rows {
		for j, cell := range row {
			if cell == '#' {
				matrix[i][j] = 1
			}
		}
	}

	s = &Shape{
		Width:  width,
		Height: height,
		Matrix: matrix,
	}
}

func (s *Shape) printShape() {
	for _, row := range s.Matrix {
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

func (s *Shape) rotateCCW() {
	// Create a new matrix with the same dimensions as the original shape but with the width and height swapped
	newMatrix := make([][]int, s.Width)
	for i := range newMatrix {
		newMatrix[i] = make([]int, s.Height)
	}

	// Fill the new matrix with the rotated values (counter-clockwise)
	for i := 0; i < s.Height; i++ {
		for j := 0; j < s.Width; j++ {
			newMatrix[s.Width-j-1][i] = s.Matrix[i][j]
		}
	}

	// Update the item's shape with the new shape
	s.Matrix = newMatrix
	s.Width, s.Height = s.Height, s.Width
}

func (s *Shape) rotateCW() {
	// Create a new matrix with the same dimensions as the original shape but with the width and height swapped
	newMatrix := make([][]int, s.Width)
	for i := range newMatrix {
		newMatrix[i] = make([]int, s.Height)
	}

	// Fill the new matrix with the rotated values (clockwise)
	for i := 0; i < s.Height; i++ {
		for j := 0; j < s.Width; j++ {
			newMatrix[j][s.Height-i-1] = s.Matrix[i][j]
		}
	}

	// Update the item's shape with the new shape
	s.Matrix = newMatrix
	s.Width, s.Height = s.Height, s.Width // Swap width and height
}

func (s *Shape) flip() {
	// rotate the item twice to get the flipped shape
	s.rotateCW()
	s.rotateCW()
}
