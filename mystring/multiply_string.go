package mystring

import (
	"fmt"
	"strconv"
)

func Multiply(s1, s2 string) string {
	m, n := len(s1), len(s2)
	ret := make([]int, m+n)
	for i := m-1; i >= 0; i-- {
		for j := n-1; j >=0; j-- {
			fmt.Println(i, j)
			a, _ := strconv.Atoi(string(s1[i]))
			b, _ := strconv.Atoi(string(s2[j]))
			product := a * b + ret[i+j+1]
			ret[i+j+1] = product % 10
			ret[i+j] += product / 10
		}
	}
	for i := 0; i < m+n; i++ {
		if ret[i] != 0 {
			return toString(ret[i:])
		}
	}
	return "0"
}

func toString(A []int) string {
	ret := ""
	for i := 0; i < len(A); i++ {
		ret += strconv.Itoa(A[i])
	}
	return ret
}


func main() {
	s1, s2 := "2", "3"
	fmt.Println(Multiply(s1, s2))
}