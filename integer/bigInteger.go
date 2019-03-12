package integer

import "fmt"
import "math/big"

func main() {
	n := big.NewInt(0)
	n.SetBit(n, 500, 1)
	n.SetBit(n, 2, 1)
	fmt.Println(n)
	fmt.Println(n.ProbablyPrime(10))
}
