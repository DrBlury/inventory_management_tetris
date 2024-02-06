package domain

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Item struct {
	ID          int    `yaml:"id"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	RawHeight   int    `yaml:"height"`
	RawWidth    int    `yaml:"width"`
	RawShape    string `yaml:"rawshape"`
	Shape       Shape
	Value       int `yaml:"value"`
}

type Shape struct {
	Width  int
	Height int
	Matrix [][]int
}

type Items struct {
	Items []*Item `yaml:"items"`
}

func parseYAML(filename string) (*Items, error) {
	var items Items

	// create reader for file
	fileHandle, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	// read the file
	fileContent, err := io.ReadAll(fileHandle)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// unmarshal the file content into the item struct
	err = yaml.Unmarshal(fileContent, &items)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling file: %w", err)
	}

	for _, item := range items.Items {
		// parse the raw shape into a matrix
		parsedShape := parseShape(item.RawShape, item.RawWidth, item.RawHeight)
		item.Shape = parsedShape
		if err != nil {
			return &items, fmt.Errorf("error parsing shape: %w", err)
		}
	}

	return &items, nil
}

func parseShape(rawShape string, width int, height int) Shape {
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

	return Shape{
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
