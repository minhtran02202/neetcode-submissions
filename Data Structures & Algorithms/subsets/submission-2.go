func subsets(nums []int) [][]int {
	var res [][]int
	var path []int

	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		// choices: add or not
		
		// add
		path = append(path, nums[i])
		backtrack(i+1)

		path = path[:len(path)-1]

		// not
		backtrack(i+1)

	}

	backtrack(0)

	return res
}
