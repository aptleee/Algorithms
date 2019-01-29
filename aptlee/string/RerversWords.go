package main

import (
	"fmt"
	"github/aptlee/mystring"
)

func ReverseWords(s string) string{
	A := mystring.Split(s, " ")
	i, j := 0, len(A)-1
	for i < j {
		A[i], A[j] = A[j], A[i]
		i++
		j--
	}
	return mystring.Join(A, " ")
}

func main() {
	s := "   the sky is blue         "
	a := mystring.Split(s, " ")
	ss := mystring.Join(a, " ")
	fmt.Println(ss, len(ss))
	fmt.Println(ReverseWords(s), len(ReverseWords(s)))
}
