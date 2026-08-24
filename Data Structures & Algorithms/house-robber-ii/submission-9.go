func rob(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	// from 0 - element before last (cause we don't want to handle last element)
	house1, house2 := 0, 0
	for i := 0; i < n - 1; i++ {
		house1, house2 = house2, max(house2, house1 + nums[i])
	}
	robsStart := house2
	
	house1, house2 = 0, 0
	for i := 1; i < n; i++ {
		house1, house2 = house2, max(house2, house1 + nums[i])
	}

	robsEnd := house2

	return max(robsStart, robsEnd)
}
