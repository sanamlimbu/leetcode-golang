package leetcode

import "math"

func MaxProfit(prices []int) int {
	i := 0 // buying day
	j := 1 // selling day
	maxProfit := 0

	for j < len(prices) {
		// profitable ?
		if prices[i] < prices[j] {
			profit := prices[j] - prices[i]
			maxProfit = int(math.Max(float64(maxProfit), float64(profit)))
		} else {
			i = j
		}
		j++
	}
	return maxProfit
}
