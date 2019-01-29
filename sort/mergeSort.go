package main

import (
	"fmt"
	"math"
)

func merge(A []int, lo, i, hi int) {
	if A[i] <= A[i+1] {
		return
	}
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

func MergeSort(A []int, lo, hi int) {
	if lo < hi {
		q := (lo+hi)/2
		MergeSort(A, lo, q)
		MergeSort(A, q+1, hi)
		merge(A, lo, q, hi)
	}
}

func MergeSortv2(A []int) {
	n := len(A)
	for sz := 1; sz < n; sz = sz * 2 {
		for i := 0; i < n-sz; i += 2*sz {
			merge(A, i, i+sz-1, min(i+sz*2-1, n-1)) //i+sz-1 for sz includes the first 1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	A := []int{9,8,7,6,5,4,3,2,1}
	MergeSortv2(A)
	fmt.Println(A)
}
