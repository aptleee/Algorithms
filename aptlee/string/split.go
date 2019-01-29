package mystring

func Split(s, e string) []string {
	ss := []string{}
	t := ""
	s = s + e
	for i := 0; i < len(s); i++ {
		if string(s[i]) != e {
			t += string(s[i])
		} else {
			if t != "" {
				ss = append(ss, t)
				t = ""
			}
		}
	}
	return ss
}
