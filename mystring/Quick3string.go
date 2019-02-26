package main


import "fmt"

func MSD(A []string, lo, hi, d int) {
	if lo >= hi {
		return
	}
	lt, i, gt := lo, lo+1, hi
	key := A[lo][d]

	for i <= gt {
		if A[i][d] < key {
			A[i], A[lt] = A[lt], A[i]
			lt++
			i++
		} else if A[i][d] < key {
				A[i], A[gt] = A[gt], A[i]
				gt--
		} else {
					i++
		}

	}
	MSD(A, lo, lt-1, d)
	if d < len(A)-1 {
		MSD(A, lt, gt, d+1)
	}
	MSD(A, gt+1, hi, d)
}

func main() {
	s := []string{"nihao", "hello", "world", "times", "66666"}
	MSD(s, 0, len(s)-1, 0)

	fmt.Println(s)
}