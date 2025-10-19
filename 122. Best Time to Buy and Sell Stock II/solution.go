package main

func maxProfit(prices []int) int {
	profit := 0
	cost := prices[0]

	for i := range prices {
		if cost < prices[i] {
			profit += prices[i] - cost
		}
		cost = prices[i]
	}

	return profit
}
