package main

import (
	"fmt"
	"strconv"
	"strings"
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

	m3 := [][]int{
		{1, -4, 9},
		{5, 2, 1},
	}

	sum := Add(m1, m2)
	subtr := Subtr(m1, m2)
	mult := Mult(m1, m2)
	transp := Transpose(m3)

	fmt.Println(Format(sum, "m"))
	fmt.Println(Format(subtr, "m"))
	fmt.Println(Format(mult, "m"))
	fmt.Println(Format(transp, "m"))
}

func NewMatrix[T Numeric](r int, c int) [][]T {
	m := make([][]T, r)

	for i := range m {
		m[i] = make([]T, c)
	}

	return m
}

func Add[T Numeric](m1 [][]T, m2 [][]T, subtract ...bool) [][]T {
	if f1, f2 := Format(m1, "ds"), Format(m2, "ds"); f1 != f2 {
		panic(fmt.Errorf("dimension mismatch: m1=%v, m2=%v", f1, f2))
	}

	sum := NewMatrix[T](len(m1), len(m1[0]))
	for i := range sum {
		for j := range sum[0] {
			if len(subtract) > 0 {
				sum[i][j] = m1[i][j] - m2[i][j]
				continue
			}

			sum[i][j] = m1[i][j] + m2[i][j]
		}
	}

	return sum
}

func Subtr[T Numeric](m1 [][]T, m2 [][]T) [][]T {
	return Add(m1, m2, true)
}

func Mult[T Numeric](m1 [][]T, m2 [][]T) [][]T {
	if !multipliable(m1, m2) {
		panic(fmt.Sprintf(
			"matrices are not multipliable: m1=%v, m2=%v\n",
			Format(m1, "ds"),
			Format(m2, "ds"),
		))
	}

	m1M, err := strconv.Atoi(strings.Split(Format(m1, "ds"), "x")[0])
	if err != nil {
		panic(err)
	}

	m2N, err := strconv.Atoi(strings.Split(Format(m2, "ds"), "x")[1])
	if err != nil {
		panic(err)
	}

	result := NewMatrix[T](m1M, m2N)
	for i := range result {
		for j := range m1[0] {
			for k := range m2[0] {
				result[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}

	return result
}

func Transpose[T Numeric](m [][]T) [][]T {
	rows, err := strconv.Atoi(strings.Split(Format(m, "ds"), "x")[0])
	if err != nil {
		panic(err)
	}

	cols, err := strconv.Atoi(strings.Split(Format(m, "ds"), "x")[1])
	if err != nil {
		panic(err)
	}

	transposed := NewMatrix[T](cols, rows)

	for i := range transposed {
		for j := range transposed[0] {
			transposed[i][j] = m[j][i]
		}
	}

	return transposed
}

func Format[T Numeric](m [][]T, f string) string {
	switch f {
	case "m":
		for _, row := range m {
			fmt.Println(row)
		}

	case "d":
		fmt.Printf("Matrix dimension: %vx%v\n", len(m), len(m[0]))

	case "ds":
		return fmt.Sprintf("%vx%v", len(m), len(m[0]))

	default:
		panic(fmt.Sprintf("'%v' is not a valid format option", f))
	}

	return ""
}

func multipliable[T Numeric](m1 [][]T, m2 [][]T) bool {
	m1N := strings.Split(Format(m1, "ds"), "x")[1]
	m2M := strings.Split(Format(m2, "ds"), "x")[0]

	return m1N == m2M
}
