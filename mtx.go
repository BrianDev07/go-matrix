package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

// TODO: mirar las operaciones más comunes con matrices e implementarlas
// suma, resta, multiplicación, división, transposición
func main() {
	size := 4
	m1 := NewMatrix[int](size, size+1)
	m2 := NewMatrix[int](size, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				m1[i][j] = 1
				m2[i][j] = 6
			}
		}
	}

	DisplayMatrix(m1)
	fmt.Println("-------------------")
	DisplayMatrix(m2)
	fmt.Println("-------------------")

	m, err := mult(m1, m2)
	if err != nil {
		panic(err)
	}

	DisplayMatrix(m)
}

// Creates a zeroed matrix of Numeric type.
func NewMatrix[T Numeric](rows int, cols int) [][]T {
	matrix := make([][]T, 0)

	for i := 0; i < rows; i++ {
		matrix = append(matrix, make([]T, cols))
	}

	return matrix
}

// Adds two matrices with the same dimension.
func Add[T Numeric](m1 [][]T, m2 [][]T, subtract ...bool) ([][]T, error) {
	if Dimension(m1) != Dimension(m2) {
		return nil, fmt.Errorf("matrix dimensions do not match: %v, %v", Dimension(m1), Dimension(m2))
	}

	res := NewMatrix[T](len(m1), len(m1))

	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m1); j++ {
			if len(subtract) == 1 {
				res[i][j] = m1[i][j] - m2[i][j]
				continue
			}

			res[i][j] = m1[i][j] + m2[i][j]
		}
	}

	return res, nil
}

// Subtracts two matrices with the same dimension.
func Subtr[T Numeric](m1 [][]T, m2 [][]T) ([][]T, error) {
	return Add(m1, m2, true)
}

// TODO: implementar multiplicación
// Multiplies two matrices only if the number of rows of the first one is equal to
// the number of columns of the second one.
func mult(m1 [][]int, m2 [][]int, div ...bool) ([][]int, error) {
	if len(m1[0]) != len(m2) {
		return nil, fmt.Errorf("matrices are not multipliable: %v, %v", Dimension(m1), Dimension(m2))
	}

	var res [][]int

	return res, nil
}

func Divide(m1 [][]int, m2 [][]int) ([][]int, error) {
	return mult(m1, m2, true)
}

// Returns a string representation of the NxM matrix dimension.
func Dimension[T Numeric](m [][]T) string {
	return fmt.Sprintf("%vx%v", len(m), len(m[0]))
}

// Shows a matrix view of the multidimensional slice.
func DisplayMatrix[T Numeric](mtx [][]T) {
	for _, row := range mtx {
		fmt.Println(row)
	}
}
