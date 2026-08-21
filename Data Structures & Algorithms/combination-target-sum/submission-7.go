import "slices"

func combinationSum(nums []int, target int) [][]int {
	slices.Sort(nums)
    var res [][]int
	var path []int
	var dfs func(int, int)

	dfs = func(candidatesStartIndex, sum int) {
		// base
		if target == sum {
			t := append([]int{}, path...)
			res = append(res, t)
		}

		// choices
		for i:=candidatesStartIndex; i<len(nums); i++ {
			// kill condition
			val := nums[i]
			if sum + val > target {
				return
			}

			// backtrack
			path = append(path, val)
			dfs(i, sum+val)
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0)

	return res
}
