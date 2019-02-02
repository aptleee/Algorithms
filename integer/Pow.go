package integer

import "fmt"

func Pow(x float32, N int) float32{
	if N < 0 {
		x = 1/x
		N = -N
	}
	return fastPowv2(x, N)
}

func fastPow(x float32, N int) float32 {
	if N == 0 {
		return 1.0
	}
	half := fastPow(x, N/2)
	if N % 2 == 0 {
		return half * half
	} else {
		return half * half * float32(x)
	}
}

func fastPowv2(x float32, N int) float32 {
	var res float32 = 1.0
	cur := x
	for i := N; i != 0; i /= 2{
		if i % 2 == 1 { // for odd numbers, the first iteration, we get res = x, and for the last iteration, where i == 1
				// res = x * cur
			res = res * cur
		}
		cur = cur * cur
	}
	return res
}

func main() {
	fmt.Println(Pow(2, -5))
}
