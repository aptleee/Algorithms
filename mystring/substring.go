package mystring

func bruteForce(s, p string) int {
	m, n := len(s), len(p)
	for i := 0; i < m-n; i++ {
		j := 0
		for j = 0; j < n; j++ {
			if s[i+j] != p[j] {
				break
			}
		}
		if j == n {
			return i
		}
	}
	return m
}

func KMP(s, p string) int {
	dfa := make([][]int, 0)
}