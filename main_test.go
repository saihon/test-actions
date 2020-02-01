package main

import "testing"

func TestFoo(t *testing.T) {
	expect := "foo"
	actual := foo()
	if actual != expect {
		t.Errorf("\ngot : %v, want: %v\n", actual, expect)
	}
}
