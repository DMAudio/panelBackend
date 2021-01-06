package typeUtils

import "testing"

func TestIndexOfStringSlice(t *testing.T) {
	var slice = []string{"a", "b", "c"}
	if IndexOfStringSlice(slice, "a") != 0 {
		t.Error("a")
	}
	if IndexOfStringSlice(slice, "b") != 1 {
		t.Error("b")
	}
	if IndexOfStringSlice(slice, "d") != -1 {
		t.Error("d")
	}
}
