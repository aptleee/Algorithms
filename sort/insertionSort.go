package insertionsort

//In the worst case scenario (reverse-sorted input), to ``insert" the i-th element, you need roughly log2i comparisons and i shifts.
import "fmt"

func Sort(A []int) {
	n := len(A)
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && A[j] > key {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
	}
}

func Sortv2(A []int) {
	n := len(A)
	for i := 1; i < n; i++ {
		key := A[i]
		j := i
		for j > 0 && A[j-1] > key {
			A[j] = A[j-1]
			j--
		}
		A[j] = key
	}
}

func searchInsert(A []int, target int) int {
	lo, hi := 0, len(A)-1
	mid := 0
	for lo <= hi {
		mid = (lo + hi) / 2
		if target > A[mid] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}
func binarySearch(A []int, target int, lo, hi int) int {
	var mid int
	for lo <= hi {
		mid = (lo + hi) / 2
		if target == A[mid] {
			return mid
		} else if target > A[mid] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

func main() {
	A := []int{5, 6, 8, 2, 6, 7, 3}
	Sort(A)
	fmt.Println(A)
	B := []int{3, 7, 9, 0, 6, 4}
	Sortv2(B)
	fmt.Println(B)
}
