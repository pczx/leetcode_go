package solution63

import "math"

func maxProfit(prices []int) int {
	maxProfit, minPrice := 0, math.MaxInt64
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price - minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}