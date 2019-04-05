package DP

func LGP(s string) int {

	var helper func(i, j int) int
	helper = func(i, j int) int {
		if i > j {
			return 0
		}
		if i == j {
			return 1
		}

		if s[i] == s[j] {
			return 2 + helper(i+1, j-1)
		}
		return max(helper(i+1, j), helper(i, j-1))
	}
	return helper(0, len(s)-1)

}

func LGP2(s string) int {
	n := len(s)

	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for l := 2; l < n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i+l-1
			if s[i] == s[j] {
				if l == 2 {
					dp[i][j] = 2
				} else {
					dp[i][j] = dp[i+1][j-1] + 2
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}

		}
	}
	
	return dp[1][n-1]
}
