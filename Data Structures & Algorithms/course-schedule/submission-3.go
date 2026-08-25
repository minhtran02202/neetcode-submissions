func canFinish(numCourses int, prerequisites [][]int) bool {
	hash := make(map[int][]int)

	for _, val := range prerequisites {
		if len(hash[val[0]]) == 0 {
			hash[val[0]] = []int{}
		}
		hash[val[0]] = append(hash[val[0]], val[1])
	}

	visited := make(map[int]bool)

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if visited[i] {
			return false
		}

		if len(hash[i]) == 0 {
			return true
		}

		visited[i] = true
		for _, val := range hash[i] {
			if !dfs(val) {
				return false
			}
		}
		visited[i] = false

		delete(hash, i)
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) { return false }
	}

	return true
}
