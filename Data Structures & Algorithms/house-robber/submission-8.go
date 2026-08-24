func rob(nums []int) int {
    house1, house2 := 0, 0
	
	for _, val := range nums {
		house1, house2 = house2, max(house2, val + house1)
	}

	return house2
}
