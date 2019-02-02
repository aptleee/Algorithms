package integer

func factorial(n int) int {
	fac := 1
	for i := 1; i <= n; i++ {
		fac *= i
	}
	return fac
}

func combanation(m , n int) int {
	return factorial(n)/(factorial(m)*factorial(n-m))
}