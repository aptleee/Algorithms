package main

import (
	"fmt"
	"math"
)

func MergeSort(A []int, lo, hi int) {
	if lo < hi {
		q := (lo+hi)/2
		MergeSort(A, lo, q)
		MergeSort(A, q+1, hi)
		merge(A, lo, q, hi)
	}
}

func merge(A []int, lo, i, hi int) {
	B := make([]int, i+2-lo)
	C := make([]int, hi-i+1)
	for j := 0; j < i-lo+1; j++ {
		B[j] = A[lo+j]
	}
	for j := 0; j < hi-i; j++ {
		C[j] = A[i+1+j]
	}
	B[i+1-lo] = math.MaxInt32
	C[hi-i] = math.MaxInt32
	m, n := 0, 0
	for k := lo; k < hi+1; k++ {
		if B[m] <= C[n] {
			A[k] = B[m]
			m++
		} else {
				A[k] = C[n]
				n++
		}
	}
}

func main() {
	A := []int{9,8,7,6,5,4,3,2,1}
	MergeSort(A, 0, len(A)-1)
	fmt.Println(A)
}
