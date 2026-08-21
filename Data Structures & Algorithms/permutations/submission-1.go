func permute(nums []int) [][]int {
	visited := make(map[int]bool)
	var res [][]int

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			t := append([]int{}, path...)
			res = append(res, t)
			return
		}	


		for i, val := range nums {
			if visited[i] {
				continue
			}

			path = append(path, val)
			visited[i] = true

			dfs(path)

			path = path[:len(path)-1]
			visited[i] = false
		}
	}

	dfs([]int{})

	return res
}
