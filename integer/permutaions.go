package integer

func permutaions(A []int) [][]int {
	ret := make([][]int, 0)
	n := len(A)
	var helper func(t []int)

	helper = func(t []int) {
		if len(t) == n {
			e := make([]int, len(t))
			copy(e, t)
			ret = append(ret, e)
		}
		for _, v := range A {
			if !contains(t, v) {
				t = append(t, v)
				helper(t)
				t = t[:len(t)-1]
			}
		}
	}
	helper([]int{})
	return ret
}


func contains(A []int, k int) bool {
	for _, i := range A {
		if i == k {
			return true
		}
	}
	return false
}
