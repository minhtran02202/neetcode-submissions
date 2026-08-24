func change(amount int, coins []int) int {
	combs := make([][]int, len(coins)+1)

	for i := range combs{
		combs[i] = make([]int, amount+1)
	}

	combs[len(coins)][0] = 1	

	for c:=len(coins)-1; c >=0; c-- {
		for a:=0; a<=amount; a++ {
			combs[c][a] = combs[c+1][a] 
			if coins[c] <= a {
				combs[c][a] += combs[c][a-coins[c]]
			}
		}
	}

	return combs[0][amount]
}
