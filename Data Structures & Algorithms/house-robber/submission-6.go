func rob(nums []int) int {
	house1, house2 := 0, 0
	
	for i := len(nums) - 1; i >= 0; i-- {
		skip := 0
		rob := nums[i]
		if i+1 < len(nums) {
			skip = house1
		}

		if i+2 < len(nums) {
			rob += house2
		}

		house2 = house1
		house1 = max(skip, rob)
	}

	return house1
}
