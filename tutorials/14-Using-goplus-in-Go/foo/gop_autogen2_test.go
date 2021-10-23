package foo_test

import (
	testing "testing"
	foo "github.com/goplus/gop/tutorial/14-Using-goplus-in-Go/foo"
)

func TestReverseMap(t *testing.T) {
	out := foo.ReverseMap(map[string]int{"b": 2})
	if len(out) != 1 || out[2] != "b" {
		t.Fatal("ReverseMap failed:", out)
	}
}
