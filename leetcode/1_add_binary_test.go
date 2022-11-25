package leetcode

import "testing"

func TestAddBinary(t *testing.T) {
	str := AddBinary("11", "1")
	if str != "100" {
		t.Errorf("Wanted 100 got %s", str)
	}

	str = AddBinary("1010", "1011")
	if str != "10101" {
		t.Errorf("Wanted 10101 got %s", str)
	}

	str = AddBinary("1110", "1011")
	if str != "11001" {
		t.Errorf("Wanted 11001 got %s", str)
	}
}
