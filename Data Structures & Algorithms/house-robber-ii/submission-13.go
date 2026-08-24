func rob(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	// from 0 - element before last (cause we don't want to handle last element)
	robsStart := handle(nums[:n-1])
	robsEnd := handle(nums[1:])

	return max(robsStart, robsEnd)
}

func handle(nums []int) int {
	house1, house2 := 0, 0
	
	for _, val := range nums {
		house1, house2 = house2, max(house2, house1 + val)
	}

	return house2 
}