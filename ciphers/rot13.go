package ciphers

func rot13(b byte) byte {
	var a byte
	switch {
	case 'a' <= b && b <= 'z':
		a = 'a'
	case 'A' <= b && b <= 'Z':
		a = 'A'
	default:
		return b
	}
	return (b-a+13)%26 + a
}


