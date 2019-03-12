package integer

import "fmt"

func findKth(A, B []int, k int) int { // given two sorted slice and find the Kth item among them
	L := len(A) + len(B)
	if L < k {
		return -1
	}
	if len(A) == 0 {
		return B[k]
	}
	if len(B) == 0 {
		return A[k]
	}

	ia, ib := len(A)/2, len(B)/2 // mid index of A and B
	ma, mb := A[ia], B[ib] // mid value of A and B

	if ia + ib < k {
		if ma > mb { // we ignore the first half of B
			return findKth(A, B[ib+1:], k-ib-1)
		}
		return findKth(A[ia+1:], B, k-ia-1)
	}
	if ma > mb { // we ignore the second half of A
		return findKth(A[:ia], B, k)
	}
	return findKth(A, B[:ib], k)
}



func Select(A []int, k int) int {
	lo, hi := 0, len(A) - 1
	for hi >= lo {
		j := partition(A, lo, hi)
		if j == k {
			return A[k]
		} else if j > k {
			hi = j - 1
		} else {
			lo = j + 1
		}
	}
	return A[k]

}

func main() {
	A := []int{1, 2, 3, 8}
	B := []int{5, 6, 7, 9}
	r := findKth(A, B, 4)
	fmt.Println(r)
}