package main

import (
	"fmt"
	"time"
)

func shellsort(A []int) {
	n := len(A)
	h := 1
	for h < n/3 {
		h = 3*h+1
	}
	for h >= 1 { // when h == 1, it's like the insertion sort
		for i := h; i < n; i++ {
			for j := i; j >= h && A[j] < A[j-h]; j -= h {
				A[j], A[j-h] = A[j-h], A[j]
			}
		}
		h = h/3
	}
}

func main() {
	a := []int{5,4,3,2,1,6,7,8,1,1}
	fmt.Println("before sort", a)
	start := time.Now()
	shellsort(a)
	elapsed := time.Since(start)
	fmt.Println("after sort", a)
	fmt.Println("shell sort took ", elapsed)
}
