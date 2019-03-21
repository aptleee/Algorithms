package DP


func knapsack(v, w []int, i, j int) int {
	if i == 0 || j == 0 {
		return 0
	}

	if j < w[i] {
		return knapsack(v, w, i-1, j)
	}
	return max(knapsack(v, w, i-1, j-w[i])+v[i], knapsack(v, w, i-1, j))
}

func knapsack1(v, w []int, i, j, W int) int {
	n := len(v)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}
	for i := 1; i <= n; i++{
		for j = 1; j <= W; j++ {
			if w[i] < j {
				dp[i][j] = max(dp[i-1][j-w[i]]+v[i], dp[i][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][W]
}