package integer


import "fmt"

func quick(A []int, k int) int {
	n := len(A)
	lo, hi := 0, n-1
	for lo <= hi {
		idx := parition(A, lo, hi)
		if idx == k {
			return A[idx]
		} else if idx > k {
				hi = idx - 1
		} else {
				lo = idx + 1
		}
	}
	return A[k]
}

func parition(A []int, lo, hi int) int {
	key := A[hi]
	pivot := lo-1
	for i := lo; i <= hi; i++ {
		if A[i] <= key {
			pivot++
			A[pivot], A[i] = A[i], A[pivot]
		}
	}
	return pivot
}

func quick3(A []int, lo, hi int) int {
	lt, i, gt := lo, lo+1, hi
	key := A[lo]
	for i <= gt {
		if A[i] < key { // lt is the idx of the first element that is equal to key
			A[lt], A[i] = A[i], A[lt]
			lt++
			i++
		} else if A[i] > key {
				A[gt], A[i] = A[i], A[gt]
				gt--
		} else {
					i++
		}
	}
	return lt
}

func main() {
	A := []int {4, 3, 9, 0, 1, 6, 3}
	fmt.Println(quick(A, 3))
}