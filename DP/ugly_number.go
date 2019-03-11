package DP

func nthUglyNumber(n int) int {

	i, j, k := 0, 0, 0

	res := make([]int, 1)
	res[0] = 1

	for len(res) < n {
		res = append(res, min(min(res[i]*2, res[j]*3), res[k]*5))
		if res[len(res)-1] == res[i]*2 {
			i++
		}
		if res[len(res)-1] == res[j]*3 {
			j++
		}
		if res[len(res)-1] == res[k]*5 {
			k++
		}
	}
	return res[len(res)-1]

}
