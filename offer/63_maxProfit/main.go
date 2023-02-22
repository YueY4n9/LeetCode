package main

func maxProfit(prices []int) int {
	profit := 0
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	for i := range prices {
		if prices[i]-min > profit {
			profit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return profit
}
