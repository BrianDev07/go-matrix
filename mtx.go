package main

import (
	"fmt"
)

// TODO: mirar las operaciones más comunes con matrices e implementarlas
// suma, resta, multiplicación, división, transposición
func main() {
	size := 9
	m1 := newMatrix[int](size, size)

	//	for i := 0; i < size; i++ {
	//		for j := 0; j < size; j++ {
	//			if i == j {
	//				matrix[i][j] = 1
	//			}
	//		}
	//	}

	displayMatrix(m1)
	fmt.Println(dimension(m1))
}

// Creates a zeroed matrix of type Type
func newMatrix[Type any](rows int, cols int) [][]Type {
	matrix := make([][]Type, 0)

	for i := 0; i < rows; i++ {
		matrix = append(matrix, make([]Type, cols))
	}

	return matrix
}

// Returns a string representation of the N*M matrix dimension
func dimension[Type any](m [][]Type) string {
	return fmt.Sprintf("%v*%v", len(m), len(m[0]))
}

// Shows a matrix view of the multidimensional slice
func displayMatrix[Type any](mtx [][]Type) {
	for _, row := range mtx {
		fmt.Println(row)
	}
}
