import "slices"

func combinationSum2(nums []int, target int) [][]int {
	slices.Sort(nums)
    var res [][]int
	var path []int
	var dfs func(int, int)

	dfs = func(start, target int) {
		if target == 0 && unique(path, res) {
			t := append([]int{}, path...)
			res = append(res, t)
			return
		}

		for i:=start; i<len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			if target - nums[i] < 0 {
				return
			}
			path = append(path, nums[i])
			dfs(i+1, target - nums[i])
			path = path[:len(path)-1]
		}
	}

	dfs(0, target)

	return res
}

func unique(path []int, res [][]int) bool {
	for _, r := range res {
		if len(path) != len(r) {
			continue
		}
		match := true
		for i, rv := range r {
			if rv != path[i] {
				match = false
				break
			}
		}
		if match {
			return false
		}
	}

	return true
}
