func maxProfit(prices []int) int {
	if len(prices) == 1 { return 0 }

	minPrice := prices[0]
	res := 0

	for i := 1; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		profit := prices[i] - minPrice
		res = max(res, profit)
	}

	return res
}
