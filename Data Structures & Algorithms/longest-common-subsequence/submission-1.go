func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

    result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)

		for j := 0; j < n; j++ {
			result[i][j] = -1
		}
	}

	var findLongestAt func(int, int) int
	findLongestAt = func(i int, j int) int {
		if i == m || j == n {
			return 0
		}

		if result[i][j] != -1 {
			return result[i][j]
		}

		if text1[i] == text2[j] {
			result[i][j] = 1 + findLongestAt(i + 1, j + 1)
		} else {
			result[i][j] = max(findLongestAt(i, j + 1), findLongestAt(i + 1, j))
		}

		return result[i][j]
	}

	return findLongestAt(0, 0)
}
