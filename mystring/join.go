package mystring

func Join(ss []string, e string) string {
	n := len(ss)
	s := ""
	for i := 0; i < n-1; i++ {
		s = s + ss[i] + e
	}
	s += ss[n-1]
	return s
}
