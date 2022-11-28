package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	profit := MaxProfit([]int{7, 1, 5, 3, 6, 4})
	if profit != 5 {
		t.Errorf("Wanted 5 got %d", profit)
	}

	profit = MaxProfit([]int{7, 6, 4, 3, 1})
	if profit != 0 {
		t.Errorf("Wanted 0 got %d", profit)
	}
}
