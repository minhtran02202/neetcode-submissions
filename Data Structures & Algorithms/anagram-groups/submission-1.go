func groupAnagrams(strs []string) [][]string {
	res := make(map[[26]int][]string)

	for _, str := range strs {
		var index [26]int
		for _, s := range str {
			index[s - 'a']++
		}

		res[index] = append(res[index], str)
	}

	var ans [][]string
	for _, val := range res {
		ans = append(ans, val)
	}

	return ans
}
