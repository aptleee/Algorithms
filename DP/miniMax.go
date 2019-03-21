package DP

func q4(A []int) int {
	n := len(A)
	memo := make([][][]int, 2)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = make([]int, n)
		}
	}

	return miniMax(A, 0, len(A)-1, true, memo)
}


func miniMax(A []int, lo, hi int, isMax bool, memo [][][]int) int {
	if hi == lo {
		return A[lo]
	}
	if isMax {
		if memo[0][lo][hi] != 0 {
			return memo[0][lo][hi]
		}
		if hi - lo == 1 {
			memo[0][lo][hi] = max(A[hi], A[lo])
			return max(A[hi], A[lo])
		}
		t := max(A[hi] + miniMax(A, lo, hi-1, !isMax, memo), A[lo] + miniMax(A, lo+1, hi, !isMax, memo))
		memo[0][lo][hi] = t
		return t
	} else {
		if memo[1][lo][hi] != 0 {
			return memo[1][lo][hi]
		}
		if hi - lo == 1 {
			memo[1][lo][hi] = min(A[hi], A[lo])
			return min(A[hi], A[lo])
		}
		t := min(miniMax(A, lo, hi-1, !isMax, memo), miniMax(A, lo+1, hi, !isMax, memo))
		memo[1][lo][hi] = t
		return t
	}
}


func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

