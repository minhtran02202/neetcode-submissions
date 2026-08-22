func combinationSum(nums []int, target int) [][]int {
    var res [][]int
	var path []int
	var dfs func(int, int)

	dfs = func(start, target int) {
		if target == 0 {
			t := append([]int{}, path...)
			res = append(res, t)
			return
		}

		for i:=start; i<len(nums); i++ {
			if target - nums[i] < 0 {
				continue
			}
			path = append(path, nums[i])
			dfs(i, target - nums[i])
			path = path[:len(path)-1]
		}
	}

	dfs(0, target)

	return res
}
