func maxProfit(prices []int) int {
	if len(prices)<2{return 0}

	left:=0
	maxProfit:=0

	for right:=1; right<len(prices); right++ {
		maxProfit = max(maxProfit, prices[right]-prices[left])

		if prices[left] >= prices[right]{ left=right }
	}

	return maxProfit
}
