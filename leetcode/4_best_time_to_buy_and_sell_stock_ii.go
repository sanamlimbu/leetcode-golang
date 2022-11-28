package leetcode

// Peak and valley problem, sum of all peaks is the required solution
func MaxProfitII(prices []int) int {
	i := 0 // buying day
	j := 1 // selling day
	maxProfit := 0

	for j < len(prices) {
		// profitable ?
		if prices[i] < prices[j] {
			maxProfit += prices[j] - prices[i]
			i++
		} else {
			i = j
		}
		j = i + 1
	}

	return maxProfit
}
