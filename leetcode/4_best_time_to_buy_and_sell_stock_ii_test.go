package leetcode

import "testing"

func TestMaxProfitII(t *testing.T) {
	profit := MaxProfitII([]int{7, 1, 5, 3, 6, 4})
	if profit != 7 {
		t.Errorf("Wanted 7 got %d", profit)
	}

	profit = MaxProfitII([]int{1, 2, 3, 4, 5})
	if profit != 4 {
		t.Errorf("Wanted 4 got %d", profit)
	}

	profit = MaxProfitII([]int{7, 6, 4, 3, 1})
	if profit != 0 {
		t.Errorf("Wanted 0 got %d", profit)
	}
}
