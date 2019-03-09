package main

import "fmt"

func LCS(A, B []int) int {
	m, n := len(A), len(B)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}

		}
	}
	return dp[m][n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	A := []int{1, 2, 3, 4, 7, 9, 10, 1111}
	B := []int{2, 4, 7, 9, 111, 10, 1111,10, 0}
	fmt.Println(LCS(A, B))
	fmt.Println("nihao")
}