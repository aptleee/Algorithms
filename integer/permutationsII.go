package integer

import "sort"

func permutaionsII(A []int) [][]int {

	n := len(A)
	ret := make([][]int, 0)
	used := make([]bool, n)
	sort.Ints(A)

	var helper func(t []int)
	helper = func(t []int) {
		if len(t) == n {
			e := make([]int, len(t))
			copy(e, t)
			ret = append(ret, e)
		}
		for i := range A {
			if used[i] || (i > 0 && A[i] == A[i-1] && !used[i-1]) {
				continue
			}
			used[i] = true
			t = append(t, A[i])
			helper(t)
			t = t[:len(t)-1]
			used[i] = false
		}
	}
	helper([]int{})
	return ret
}

