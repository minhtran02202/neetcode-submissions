func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	return max(nums[0], max(handle(nums[:len(nums) - 1]), handle(nums[1:])))
}

func handle(nums []int) int {
	house1, house2 := 0, 0
	
	for _, val := range nums {
		house1, house2 = house2, max(house2, house1 + val)
	}

	return house2 
}