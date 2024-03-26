package domain

import "fmt"

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
