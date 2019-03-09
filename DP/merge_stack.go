package DP


type T struct {
	i int
	j int
}

func q3(A []int) int {
	n := len(A)
	m := make(map[T]int)
	val := make([][]int, n)
	for i := 0; i < n; i++ {
		val[i] = make([]int, n)

	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val[i][j] = sum(A, i, j)
		}
	}

	var merge func(i, j int) int
	merge = func(i, j int) int {
		if v, ok := m[T{i, j }]; ok {
			return v
		}
		Min := (1 << 31) - 1
		if i == j {
			return 0
		}
		if i == j-1 {
			return A[i]+A[j]
		}
		for k := i; k <= j; k++ {
			cur := 0
			if k == i {
				cur = merge(k+1, j) + val[k][j]
			} else if k == j {
				cur = merge(i, k-1) + val[i][j]
			} else {
				cur = val[i][j] + merge(i, k)+merge(k+1, j)
			}
			Min = min(Min, cur)
		}
		m[T{i, j}] = Min
		return Min
	}
	return merge(0, n-1)
}


func sum(A []int, i, j int) int {
	ans := 0
	for k := i; k <= j; k++ {
		ans += A[k]
	}
	return ans
}

