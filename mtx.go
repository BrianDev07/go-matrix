package main

import (
	"fmt"
)

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

func main() {
	m1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	m2 := [][]int{
		{1, -4, 9},
		{5, 2, 1},
		{9, -7, -5},
	}
}

// Creates matrix with dimensions rxc and generic type T
func NewMatrix[T Numeric](r int, c int) [][]T {
	m := make([][]T, r)

	for i := range m {
		m[i] = make([]T, c)
	}

	return m
}

// Adds two matrices with the same dimension.
func Add[T Numeric](m1 [][]T, m2 [][]T, subtract ...bool) ([][]T, error) {
	if Dimension(m1) != Dimension(m2) {
		return nil, fmt.Errorf("matrix dimensions do not match: %v, %v", Dimension(m1), Dimension(m2))
	}

	res := NewMatrix[T](len(m1), len(m1)) // matriz cuadrada??????

	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m1); j++ {
			if subtract[0] == true {
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
