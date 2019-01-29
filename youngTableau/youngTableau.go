package main

import (
	"fmt"
	"math"
)

func ExtractMin(A [][]int) int {
	M := A[0][0]
	i, j := 0, 0
	m, n := len(A), len(A[0])

	for i < m && j < n {
		a, b := i, j // original coordinate
		// initialization
		right, down := math.MaxInt32, math.MaxInt32
		if i < m-1 { // when i == m-1, down == MAX, last row
			down = A[i+1][j]
		}
		if j < n-1 { // when j = n-1, right = MAX, last column
			right = A[i][j+1]
		}

		// compare right with down to decide how to swap
		if right == math.MaxInt32 && down == math.MaxInt32 { // biggest one
			A[i][j] = math.MaxInt32
			break
		} else if right < down {
			j++
		} else {
				i++
		}
		A[i][j], A[a][b] = A[a][b], A[i][j]
	}

	return M
}

func ExtractMinv2(A [][]int, p, q int) int {
	M := A[p][q]
	m, n := len(A), len(A[0])
	right, down := math.MaxInt32, math.MaxInt32
	if p < m-1 {
		down = A[p+1][q]
	}
	if q < n-1 {
		right = A[p][q+1]
	}
	if right == math.MaxInt32 && down == math.MaxInt32 { // biggest one
		A[p][q] = math.MaxInt32
	} else if right < down {
		a := ExtractMinv2(A, p, q+1)
		A[p][q] = a
	} else {
		b := ExtractMinv2(A, p+1, q)
		A[p][q] = b
	}
	return M
}

func Exist(A [][]int, k int) bool {
	m, n := len(A), len(A[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		if A[i][j] == k {
			return true
		} else if A[i][j] > k {
			j--
		} else {
				i++
		}
	}
	return false
}

func Insert(A [][]int, v int) {
	i, j := len(A)-1, len(A[0])-1
	A[i][j] = v
	var left, up int
	for i >= 0 && j >= 0 {
		left, up = math.MinInt32, math.MinInt32
		if i > 0 {
			up = A[i-1][j]
		}
		if j > 0 {
			left = A[i][j-1]
		}
		if v >= left && v >= up {
			break
		} else if v < left {
			A[i][j], A[i][j-1] = A[i][j-1], A[i][j]
			j--
		} else {
				A[i][j], A[i-1][j] = A[i-1][j], A[i][j]
				i--
		}
	}
}

func Insertv2(A [][]int, v int) {
	i, j := len(A)-1, len(A[0])-1
	A[i][j] = v
	var left, up int
	for i >= 0 && j >= 0 {
		left, up = math.MinInt32, math.MinInt32
		if i > 0 {
			up = A[i-1][j]
		}
		if j > 0 {
			left = A[i][j-1]
		}
		if v >= left && v >= up {
			break
		} else if v < up {
			A[i][j], A[i-1][j] = A[i-1][j], A[i][j]
			i--
		} else {
			A[i][j], A[i][j-1] = A[i][j-1], A[i][j]
			j--
		}
	}
}


func main() {
	a := [][]int{{2, 3, 12, math.MaxInt32}, {4, 8, 16, math.MaxInt32}, {5, 9, math.MaxInt32, math.MaxInt32}, {14, math.MaxInt32, math.MaxInt32, math.MaxInt32}}
	for i := 0; i < len(a); i++{
		for j := 0; j < len(a[0]); j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
	for i := 0; i < 12; i++ {
		fmt.Print(ExtractMin(a), " ")
	}
	fmt.Println()
	for i := 0; i < 12; i++ {
		Insertv2(a, i)
	}
	for i := 0; i < len(a); i++{
		for j := 0; j < len(a[0]); j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println(Exist(a, 8))
}
