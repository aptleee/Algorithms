package integer

import (
	"fmt"
	"math"
)

func Reverse(v int) int {
	var res = 0
	for v != 0 {
		res = v % 10 + res * 10
		v = v/10
	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	} else {
		return res
	}
}

func main() {
	fmt.Println(Reverse(10))
	fmt.Println(Reverse(2147483648))
	fmt.Println(Reverse(-123))
}