import "slices"

func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	var res [][]int

	var backtrack func(i int, path []int)
	backtrack = func(i int, path []int) {
		if i == len(nums) {
			t := append([]int{}, path...)
			res = append(res, t)
			return
		}

		path = append(path, nums[i])
		backtrack(i+1, path)
		path = path[:len(path)-1]

		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
		backtrack(i+1, path)
	}

	backtrack(0, []int{})

	return res
}
