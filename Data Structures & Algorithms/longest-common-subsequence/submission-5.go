func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

    result := make([]int, n + 1)

	for i := m - 1; i >= 0; i-- {
		diag := 0
		for j := n - 1; j >= 0; j-- {
			temp := result[j]
			if text1[i] == text2[j] {
				result[j] = 1 + diag
			} else {
				result[j] = max(result[j + 1], result[j])
			}
			diag = temp
		}
	}

	return result[0]
}
