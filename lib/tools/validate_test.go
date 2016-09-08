package tools

import (
	"fmt"
	"testing"
)

func TestStringInSlice(t *testing.T) {
	cases := []struct {
		slice  []string
		target string
		want   bool
	}{
		{
			[]string{"Testing", "12345", "@#¤%&"},
			"Testing",
			true,
		},
		{
			[]string{"Testing", "12345", "@#¤%&"},
			"12345",
			true,
		},
		{
			[]string{"Testing", "12345", "@#¤%&"},
			"@#¤%&",
			true,
		},
		{
			[]string{"Testing", "12345", "@#¤%&"},
			"Testin",
			false,
		},
		{
			[]string{"Testing", "12345", "@#¤%&"},
			"&",
			false,
		},
		{
			[]string{"Testing", "12345", "@#¤%&"},
			"@#¤%",
			false,
		},
	}

	for _, c := range cases {
		got := StringInSlice(c.target, c.slice)
		if got != c.want {
			fmt.Errorf("StringInSlice(%q) == %q, want %q", c.target, got, c.want)
		}
	}
}
