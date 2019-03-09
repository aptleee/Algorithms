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

