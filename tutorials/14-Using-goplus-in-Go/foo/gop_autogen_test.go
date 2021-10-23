package foo

import testing "testing"

func TestReverseMap(t *testing.T) {
	out := ReverseMap(map[string]int{"a": 1})
	if len(out) != 1 || out[1] != "a" {
		t.Fatal("ReverseMap failed:", out)
	}
}
