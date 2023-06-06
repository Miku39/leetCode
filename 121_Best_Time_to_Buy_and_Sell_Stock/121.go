package main

// 121. Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/?envType=study-plan&id=level-1
func maxProfit(prices []int) int {
	maxProfit := 0
	lowPrice := 10000
	for _, v := range prices {
		if v < lowPrice {
			lowPrice = v
		} else if (v - lowPrice) > maxProfit {
			maxProfit = v - lowPrice
		}
	}
	return maxProfit
}
