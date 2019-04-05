package main
import "fmt"

type t struct {
	i int
	j int
}

func mergeStack1(A []int) int {
	n := len(A)
	m := make(map[t]int)
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
		if v, ok := m[t{i, j }]; ok {
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
		m[t{i, j}] = Min
		return Min
	}
	return merge(0, n-1)
}

// Algorithm:
// dp[i][j] = val[i][j] + dp[i][k] + dp[k+1][j], k from i to j-1

func mergeStack(A []int) int {
	n := len(A)
	val := make([][]int, n)
	dp := make([][]int, n)
	for i := range val {
		val[i] = make([]int, n)
		dp[i] = make([]int, n)
		for j := range val[i] {
			val[i][j] = sum(A, i, j)
			if i == j {
				dp[i][j] = 0
			} else {
				dp[i][j] = 1 << 30
			}
			
		}
	}
	for l := 2; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], val[i][j] + dp[i][k] + dp[k+1][j])
			}
		}
	}
	return dp[0][n-1]
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func main() {
	A := []int{1,2,3,4,5,1,2,3,2,3,41,2222,1333,21,3,42,2,4,6}
	fmt.Println(mergeStack(A), mergeStack1(A))
}


func sum(A []int, i, j int) int {
	ans := 0
	for k := i; k <= j; k++ {
		ans += A[k]
	}
	return ans
}

