package main

import (
	"testing"
	"unicode/utf8"
)

// func TestReverse(t *testing.T) {
// 	testcases := []struct {
// 		in, want string
// 	}{
// 		{"Hello, world", "dlrow ,olleH"},
// 		{" ", " "},
// 		{"!12345", "54321!"},
// 	}

// 	for _, tc := range testcases {
// 		rev := Reverse(tc.in)
// 		if rev != tc.want {
// 			t.Errorf("Reverse: %q, want %q", rev, tc.want)
// 		}
// 	}
// }

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, original string) {
		rev, err1 := Reverse(original)
		if err1 != nil {
			return
		}

		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}

		if original != doubleRev {
			t.Errorf("Before: %q, after %q", original, doubleRev)
		}
		if utf8.ValidString(original) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
