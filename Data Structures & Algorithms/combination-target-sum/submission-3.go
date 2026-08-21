func combinationSum(nums []int, target int) [][]int {
    var res [][]int

	var dfs func([]int, []int, int)

	dfs = func(candidates, path []int, sum int) {
		// base
		if target == sum {
			t := append([]int{}, path...)
			res = append(res, t)
		}

		// choices
		for i, val := range candidates {
			// kill condition
			if sum + val > target {
				continue
			}

			// backtrack
			path = append(path, val)
			dfs(candidates[i:], path, sum+val)
			path = path[:len(path)-1]
		}
	}

	dfs(nums, []int{}, 0)

	return res
}
