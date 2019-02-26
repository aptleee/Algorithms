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
	// when there is a match, so next position of j is j+1, we get dfa[][j] = j+1
	// when there is a mismatch, we know not just s[i],
	// but also the j-1 previous characters in the text: they are the first j-1 characters in the pattern.
	// so we slide a copy of p on the j characters, and move from left to right, until there is a match
	// in this way we can find the next position of p to be compared with next character in s
	//
	R, m := 256, len(p)
	dfa := make([][]int, R) // dfa[c][j] is the position of p to compare against next s , after compare c with p[j]
	for i := 0; i < R; i++ {
		dfa[i] = make([]int, m)
	}

	dfa[p[0]][0] = 1
	for x, j := 0, 1; j < m; j++ {
		for c := 0; c < R; c++ {
			dfa[c][j] = dfa[c][x]
		}
		dfa[p[j]][j] = j+1
		x = dfa[p[j]][x]
	}

	search := func(s string) int {
		i, j, n, m := 0, 0, len(s), len(p)
		for ; i < n && j < m; i++ {
			j = dfa[s[i]][j] // move j according to dfa[][j]
		}
		if j == m {
			return i - m
		}
		return n
	}
	return search(s)
}