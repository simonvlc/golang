package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		input, output string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.input)
		if got != c.output {
			t.Errorf("Reverse(%q) == %q, output %q", c.input, got, c.output)
		}
	}
}
