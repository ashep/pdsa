package main

import "testing"

var s1 = "1a2B3c4D5e6F"
var s2 = "F6e5D4c3B2a1"

func TestReverse(t *testing.T) {
	rs := reverse(s1)
	if rs != s2 {
		t.Logf("%q expected, got %q", s2, rs)
		t.Fail()
	}
}
