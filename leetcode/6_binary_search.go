package leetcode

/*
Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.
You must write an algorithm with O(log n) runtime complexity.
*/
func Search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	if len(nums) == 0 {
		return -1
	}

	for l <= r {
		mid := (l + r) / 2
		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
