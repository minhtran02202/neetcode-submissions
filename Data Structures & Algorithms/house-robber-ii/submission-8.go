func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	house1, house2 := 0, 0
	for i := 0; i < len(nums) - 1; i++ {
		house1, house2 = house2, max(house2, house1 + nums[i])
	}

	robsStart := house2

	house1, house2 = 0, 0
	for i := 1; i < len(nums); i++ {
		house1, house2 = house2, max(house2, house1 + nums[i])
	}

	robsEnd := house2

	return max(robsStart, robsEnd)
}
