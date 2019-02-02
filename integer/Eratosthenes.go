package main

import "fmt"

//given an integer n, construct a prime table range from 2 - n

func Eratosthenes(n int) []int { // delete 2p, 3p, 4p,... for every non-negative p less than n
	vis := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if vis[i] != 1 {
			for j := 2 * i; j <= n; j += i {
				vis[j] = 1
			}
		}
	}
	ret := []int{}
	for i := 2; i <= n; i++ {
		if vis[i] == 0 {
			ret = append(ret, i)
		}
	}
	return ret
}

func main() {
	fmt.Println(Eratosthenes(100))
}