package integer

import (
	"fmt"
	"math"
)

func isPrime(n int) bool { // i * i may overflow
	if n == 1 {
		return false
	}
	m := int(math.Sqrt(float64(n))+0.5) // in float, x+1 may be x.9999, after int it will be x
	// or use func iSqrt(n int) int
	for i := 2; i <= m; i++ { // if i * i > n, and when n % i == 0, n/i will less than i, which is impossible
		if n % i == 0 {
			return  false
		}
	}
	return true
}

func main() {
	fmt.Println(isPrime(8))
	fmt.Println(isPrime(3))
	fmt.Println(isPrime(11))
	fmt.Println(isPrime(19))

}