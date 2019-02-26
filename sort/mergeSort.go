package main

import (
	"fmt"
	"math"
)

func merge(A []int, lo, i, hi int) int {
	B := make([]int, i+2-lo)
	C := make([]int, hi-i+1)
	count := 0
	for j := 0; j < i-lo+1; j++ {
		B[j] = A[lo+j]
	}
	for j := 0; j < hi-i; j++ {
		C[j] = A[i+1+j]
	}
	B[i+1-lo] = math.MaxInt32
	C[hi-i] = math.MaxInt32
	m, n := 0, 0
	for k := lo; k <= hi; k++ {
		if B[m] <= C[n] {
			A[k] = B[m]
			m++
		} else {
				A[k] = C[n]
				n++
				count += (i+1-lo)-m
		}
	}
	return count
}

func Merge(A []int, lo, i, hi int) int {
	B := make([]int, i-lo+1)
	C := make([]int, hi-i)
	count := 0
	copy(B, A[lo:i+1])
	copy(C, A[i+1:hi+1])
	B = append(B, math.MaxInt32)
	C = append(C, math.MaxInt32)
	m, n := 0, 0
	for k := lo; k <= hi; k++ {
		if B[m] <= C[n] {
			A[k] = B[m]
			m++
		} else {
			A[k] = C[n]
			n++
			count += (i+1-lo)-m
		}
	}
	return count
}


func MergeSort(A []int, lo, hi int) {
	if lo < hi {
		q := (lo+hi)/2
		MergeSort(A, lo, q)
		MergeSort(A, q+1, hi)
		if A[q] <= A[q+1] {
			return
		}
		merge(A, lo, q, hi)
	}
}

func MergeSort2(A []int) {
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




func inversionCount(A []int, lo, hi int) int {
	if lo < hi {
		q := (lo+hi)/2
		c1 := inversionCount(A, lo, q)
		c2 := inversionCount(A, q+1, hi)
		c3 := Merge(A, lo, q, hi)
		return c1+c2+c3
	}
	return 0
}

func main() {
	A := []int{5, 2, 3, 4, 5}
	fmt.Println(A, inversionCount(A, 0, len(A)-1))
}
