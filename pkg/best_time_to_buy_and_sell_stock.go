package pkg

// 121. Best Time to Buy and Sell Stock
//
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock
func maxProfit(prices []int) int {
	low, high, max := 0, 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[low] {
			low = i
		}
		if prices[i] > prices[high] {
			high = i
		}
		if low < high {
			delta := prices[high] - prices[low]
			if delta > max {
				max = delta
			}
		} else {
			high = low
		}
	}

	return max
}
