package leetcode

import "testing"

func TestSearch(t *testing.T) {
	arr := []int{-1, 0, 3, 5, 9, 12}
	got := Search(arr, 9)
	if got != 4 {
		t.Errorf("wanted 4 got %d", got)
	}

	arr = []int{-1, 0, 3, 5, 9, 12}
	got = Search(arr, 2)
	if got != -1 {
		t.Errorf("wanted -1 got %d", got)
	}
}
