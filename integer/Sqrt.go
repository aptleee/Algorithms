package integer

import (
	"fmt"
)

func Sqrt(x float64) float64 { // Newton's method
	// starting with some guess z, z could be some value, such as 1 or 2 or x/2 or x
	z := float64(x/2)
	for i := 0; i < 10; i++ {
		z -= (z*z- x)/(2*z)
	}
	return z
}

func Sqrtv2(x float64) float64 {
	z := float64(x/2)
	for (z*z - x) / (2*z) > 1e-10 {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func iSqrt(x int) int {
	if x == 0 {
		return 0
	}
	lo, hi := 1, x
	var ans, mid int
	for lo <= hi {
		mid = (lo+hi)/2
		if mid <= x/mid {
			lo = mid+1
			ans = mid
		} else {
			hi = mid - 1
		}
	}
	return ans
}

func main() {
	fmt.Println(iSqrt(2147395599))
}
