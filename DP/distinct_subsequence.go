package DP


func numDistinct1(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}


	var helper func(s, t string) int
	helper = func(s, t string) int{
		if dp[len(s)][len(t)] != -1 {
			return dp[len(s)][len(t)]
		}
		res := 0
		if len(s) < len(t) {
			return 0
		}
		if s == t || len(t) == 0{
			return 1
		}

		for i := len(s)-1; i >= 0; i-- {
			if len(t) > 0 && i >= len(t)-1 && s[i] == t[len(t)-1] {
				res += numDistinct(s[:i], t[:len(t)-1])
			}
		}
		dp[len(s)][len(t)] = res
		return res
	}
	return helper(s, t)

}

func numDistinct(s string, t string) int {
	// dp[i][j] denotes how many times s[:j] contains t[:i]
	// if s[i] == s[j], dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
	// if s[i] != s]j], dp[i][j] = dp[i][j-1]
	m, n := len(s), len(t)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[0][j] = 1
		}
	}

	for i := 1; i < n+1; i++ {
		for j := i; j < m+1; j++ {
			if s[j-1] == t[i-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n]

}