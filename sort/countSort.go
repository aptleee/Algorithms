package main

import (
	"fmt"
	"math"
)

func count(A []int) {
	n := len(A)
	min, max := math.MaxInt32, math.MinInt32

	for _, x := range A {
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}

	count := make([]int, max-min+1)
	for _, x := range A {
		count[x-min]++
	}

	for i := 0; i < len(count)-1; i++ {
		count[i+1] += count[i]
	}

	aux := make([]int, n)
	for _, x := range A {
		aux[count[x-min]-1] = x
		count[x-min]--
	}
	for i := range A {
		A[i] = aux[i]
	}
}

func main() {
	A := []int{2, 3, 1, 6, 0, 8, 7}
	count(A)
	fmt.Println(A)
}