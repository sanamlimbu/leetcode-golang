package leetcode

import "testing"

func TestAddStrings(t *testing.T) {
	var sum string
	sum = AddStrings("11", "123")
	if sum != "134" {
		t.Errorf("Wanted 134 got %s", sum)
	}

	sum = AddStrings("456", "77")
	if sum != "533" {
		t.Errorf("Wanted 533 got %s", sum)
	}

	sum = AddStrings("0", "0")
	if sum != "0" {
		t.Errorf("Wanted 0 got %s", sum)
	}

	sum = AddStrings("999", "999")
	if sum != "1998" {
		t.Errorf("Wanted 1998 got %s", sum)
	}
}
