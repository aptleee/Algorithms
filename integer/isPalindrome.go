package integer

import "fmt"

func IsPalindrome(x int) bool {
	if x < 0 || x != 0 && x % 10 == 0 {
		return false
	}
	sum := 0
	for x > sum {
		sum = sum*10 + x%10
		x = x/10
	}
	return x == sum || x == sum/10 // for 121, x = 1, sum = 12 compare x and sum/10
	// for 1221, x = 12, sum = 12, compare x and sum
}

func main() {
	fmt.Println(IsPalindrome(10))
	fmt.Println(IsPalindrome(121))
	fmt.Println(IsPalindrome(12233221))
}
