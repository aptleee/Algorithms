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

func me(A, left, right []int) int {
	n1 := len(left)
	n2 := len(right)
	m, n, count := 0, 0, 0
	for m < n1 || n < n2 {
		if m == n1 {
			A[m+n] = right[n]
			n++
		} else if n == n2 {
				A[m+n] = left[m]
				m++
		} else if left[m] <= right[n] {
					A[m+n] = left[m]
					m++
		} else {
						A[m+n] = right[n]
						n++
						count += n1 - m
		}
	}
	return count
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




func inversionCount(A []int) int {
	if len(A) >= 2 {
		q := (len(A)-1)/2
		B := make([]int, q+1)
		C := make([]int, len(A)-q-1)
		copy(B, A[:q+1])
		copy(C, A[q+1:])
		a := inversionCount(B)
		b := inversionCount(C)
		c := me(A, B, C)
		return a+b+c
	}
	return 0
}

func main() {
	A := []int{5, 2, 3, 4, 5}
	fmt.Println(A, inversionCount(A))
}
