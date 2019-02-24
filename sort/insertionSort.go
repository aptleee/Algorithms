package sort


//In the worst case scenario (reverse-sorted input), to ``insert" the i-th element, you need roughly log2i comparisons and i shifts.

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

func rank(A []int, target int) int { // the same as searchInsert, which
	lo, hi := 0, len(A)-1
	mid := 0
	for lo <= hi {
		mid = (lo + hi) / 2
		if target > A[mid] {
			lo = mid + 1
		} else if target < A[mid] {
			hi = mid - 1
		} else {
			return mid
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

func search(A []int) int { // search the idx of the smallest one in the rotated array
	lo, hi := 0, len(A)-1
	for lo < hi { // when lo == hi break the loop
		mid := (lo + hi) / 2
		if A[mid] > A[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

//	root := lo
//	lo, hi = 0, len(A)-1
//	for lo <= hi {
//		mid := (lo + hi) / 2
//		realMid := (mid + root) % len(A)
//		if A[realMid] == target {
//			return realMid
//		} else if target > A[realMid] {
//			lo = mid + 1
//		} else {
//			hi = mid - 1
//		}
//	}
//	return -1
//}



func main() {

	//Sort(A)
	//fmt.Println(A)
	//B := []int{3, 7, 9, 0, 6, 4}
	//Sortv2(B)
	//fmt.Println(B)
}
