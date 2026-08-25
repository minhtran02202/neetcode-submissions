func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

    result := make([][]int, m + 1)
	for i := 0; i <= m; i++ {
		result[i] = make([]int, n + 1)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				result[i][j] = 1 + result[i + 1][j + 1]
			} else {
				result[i][j] = max(result[i][j + 1], result[i + 1][j])
			}
		}
	}

	return result[0][0]
}
