package main


import "fmt"

func LSD(A []string, W int) {
	n := len(A)
	R := 256
	aux := make([]string, n)

	for d := W-1; d >= 0; d-- {
		count := make([]int, R+1)
		for _, x := range A {
			count[x[d]]++
		}

		for i := 0; i < len(count)-1; i++ {
			count[i+1] += count[i]
		}

		for i := len(A)-1; i >= 0; i-- {
			aux[count[A[i][d]]-1] = A[i]
			count[A[i][d]]--
		}

		for i := 0; i < len(A); i++ {
			A[i] = aux[i]
		}
	}
}

func main() {
	s := []string{"nihao", "hello", "world", "times", "66666"}
	LSD(s, len(s))

	fmt.Println(s)

}