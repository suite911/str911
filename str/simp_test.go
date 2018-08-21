package main

import "testing"

func TestSimp(t *testing.T) {
	for _, x := range [][2]string{
		[2]string{"hello", "hello"},
		[2]string{"123", "123"},
		[2]string{"he@#!llo", "hello"},
		[2]string{"", ""},
		[2]string{"@#!", ""},
	} {
		if simp := Simp(x[0]); simp != x[1] {
			t.Errorf("Simp(%q) != %q // %q", x[0], x[1], simp)
		}
	}
}
