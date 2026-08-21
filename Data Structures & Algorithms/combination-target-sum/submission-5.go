import "slices"

func combinationSum(nums []int, target int) [][]int {
	slices.Sort(nums)
    var res [][]int
	var path []int
	var dfs func([]int, int)

	dfs = func(candidates []int, sum int) {
		// base
		if target == sum {
			t := append([]int{}, path...)
			res = append(res, t)
		}

		// choices
		for i, val := range candidates {
			// kill condition
			if sum + val > target {
				return
			}

			// backtrack
			path = append(path, val)
			dfs(candidates[i:], sum+val)
			path = path[:len(path)-1]
		}
	}

	dfs(nums, 0)

	return res
}
