package main

import "fmt"

func partition(A []int, lo, hi int) int{
	k := A[hi]
	i := lo - 1
	for j := lo; j < hi; j++ { // A[lo, i) <= k, A[i, j) > k, A[j, hi] to be checked
		if A[j] <= k {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[hi] = A[hi], A[i+1] // A[hi] is equal to k, A[i+1] is larger than k, so swap them
	return i+1
}

func Quick3Way(A []int, lo, hi int) int{
	k := A[lo]
	lt, i, gt := lo, lo+1, hi // A[lo-lt) < k, A[lt-i) == k, A(gt-hi] > k, A[i-gt] to be checked
	for i <= gt {
		if A[i] < k { // lt is the first one who is equals to k, A[lt] must be equal k, so after swap, A[i] == A[lt] == k, we increment lt and i
			A[lt], A[i] = A[i], A[lt]
			lt++
			i++
		} else if A[i] > k { // for we don't know whether A[gt] is greater or less than or equal to k after swap
							// so we don't move i, and check A[i] (A[gt]) in the next round
				A[gt], A[i] = A[i], A[gt]
				gt--
		} else {
					i++
		}
	}
	return lt
}

func Quick3Wayv2(A []int, lo, hi, pivot int) int {
	k := A[pivot]
	lt, i, gt := lo, lo, hi
	for i <= gt {
		if A[i] < k { // A[lt] must be equal k, so after swap, A[i] == A[lt] == k, we increment lt and i
			A[lt], A[i] = A[i], A[lt]
			lt++
			i++
		} else if A[i] > k { // for we don't know whether A[gt] is greater or less than or equal to k after swap
			// so we don't move i, and check A[i] (A[gt]) in the next round
			A[gt], A[i] = A[i], A[gt]
			gt--
		} else {
			i++
		}
	}
	return lt
}

func QuickSort(A []int, lo, hi int) {
	if lo < hi {
		q := partition(A, lo, hi)
		QuickSort(A, lo, q-1)
		QuickSort(A, q+1, hi)
	}
}

func QuickSortv2(A []int, lo, hi int) {
	if lo < hi {
		q := Quick3Way(A, lo, hi)
		QuickSort(A, lo, q-1)
		QuickSort(A, q+1, hi)
	}
}
		
func main() {
	A := []int{1, 2, 5, 2, 3, 5, 7, 1, 9}
	QuickSort(A, 0, len(A)-1)
	fmt.Println(A)
	A = []int{1, 2, 5, 2, 3, 5, 7, 1, 9}
	QuickSortv2(A, 0, len(A)-1)
	fmt.Println(A)
}
