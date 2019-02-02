package mystring

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{"世界", "界世"},
		{"hello", "olleh"},
		{"", ""},
	}
	for _, c := range tests {
		got := Reverse(c.s)
		if got != c.want {
			t.Errorf("Reverse(%q) = %q, want %q", c.s, got, c.want)
		}
	}
}
