package leetcode

import "testing"

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := MaxSubArray(nums)
	if result != 6 {
		t.Errorf("wanted 6 , got %d", result)
	}

	nums = []int{1}
	result = MaxSubArray(nums)
	if result != 1 {
		t.Errorf("wanted 1, got %d", result)
	}

	nums = []int{5, 4, -1, 7, 8}
	result = MaxSubArray(nums)
	if result != 23 {
		t.Errorf("wanted 23, got %d", result)
	}
}
