package mysort

import "fmt"

func sink(A []int, i, n int) {
	largest := i
	for {

		j := 2*i + 1
		if j > n || j < 0{
			break
		}
		if A[j] > A[i] {
			largest = j
		}
		j++
		if j <= n && A[j] > A[largest] {
			largest = j
		}
		if largest == i {
			break
		}
		A[i], A[largest] = A[largest], A[i]
		i = largest
	}
}

func heapSort(A []int) {
	n := len(A) - 1

	for i := n/2; i >= 0; i-- {
		sink(A, i, n)
	}
	for i := n; i >= 0; i-- {
		A[i], A[0] = A[0], A[i]
		sink(A, 0, i-1)
	}
}

func main() {
	A := []int{1, 3, 1, 8, 4, 5}
	heapSort(A)
	fmt.Println(A)
}