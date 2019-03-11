package DP


func minDistance2(word1 string, word2 string) int {
	// four situations
	// word1[0] == word2[0], so call minDistance(word1[1:], word2[1:])
	// word1[0] != word2[0], change word1[0] to word2[0], count++, and call minDistance(word[1:], word2[1:])
	// omit word1[0] count++, call minDistance(word[1:], word2)
	// insert word2[0] at the first of word1, count++, then call minDistance(word1, word2[1:])

	if word1 == word2 {
		return 0
	}

	var helper func(word1, word2 string) int
	helper = func(word1, word2 string) int {
		if word1 == "" {
			return len(word2)
		}
		if word2 == "" {
			return len(word1)
		}

		if word1[0] == word2[0] {
			return helper(word1[1:], word2[1:])
		}
		c2 := helper(word1[1:], word2[1:]) + 1
		c3 := helper(word1[1:], word2) + 1
		c4 := helper(word1, word2[1:]) + 1
		return min(min(c2, c3), c4)
	}

	return helper(word1, word2)
}


func minDistance3(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 {
				dp[0][j] = j
			} else if j == 0 {
				dp[i][0] = i
			} else if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j-1]+1, dp[i][j-1]+1), dp[i-1][j]+1)
			}
		}
	}
	return dp[m][n]
}

func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}
	m, n := len(word1), len(word2)
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 {
				dp[0][j] = j
			} else if j == 0 {
				dp[1][0] = i
			} else if word1[i-1] == word2[j-1] {
				dp[1][j] = dp[0][j-1]
			} else {
				dp[1][j] = min(min(dp[0][j-1]+1, dp[1][j-1]+1), dp[0][j]+1)
			}
		}
		if i > 0 {
			for k := 0; k < n+1; k++ {
				dp[0][k] = dp[1][k]
			}
		}
	}
	if m == 0 {
		return dp[0][n]
	}
	return dp[1][n]
}
