package main

func maxProfit(prices []int) int {
	profit := 0
	buy := prices[0]
	for _, day := range prices[1:] {
		if buy > day {
			buy = day
		} else if profit < day-buy {
			profit = day - buy
		}

	}
	return profit
}
