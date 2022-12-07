package leetcode

/*
	Given an integer array nums, find the subarray
 	which has the largest sum and return its sum.
*/

// MaxSubArray finds the sum of maximum subarray
func MaxSubArray(nums []int) int {
	right := len(nums) - 1
	return finMaxSubArray(nums, 0, right)
}

func finMaxSubArray(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	mid := (left + right) / 2
	leftMax := finMaxSubArray(nums, left, mid)
	rightMax := finMaxSubArray(nums, mid+1, right)
	crossMax := maxCrossing(nums, left, right, mid)
	return max(leftMax, rightMax, crossMax)
}

func maxCrossing(nums []int, left, right, mid int) int {
	leftSum := -1 << 31 // -2,147,483,648 lowest possible int value
	rightSum := -1 << 31

	sum := 0
	for i := mid; i >= left; i-- {
		sum += nums[i]
		leftSum = max(leftSum, sum)
	}

	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		rightSum = max(rightSum, sum)
	}

	return leftSum + rightSum
}

func max(values ...int) int {
	maxVal := values[0]
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
