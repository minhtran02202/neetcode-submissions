func letterCombinations(digits string) []string {
	if len(digits) == 0 {return []string{}}

	table := make(map[int][]string)
	table[2] = []string{"a", "b", "c"}
	table[3] = []string{"d", "e", "f"}
	table[4] = []string{"g", "h", "i"}
	table[5] = []string{"j", "k", "l"}
	table[6] = []string{"m", "n", "o"}
	table[7] = []string{"p", "q", "r", "s"}
	table[8] = []string{"t", "u", "v"}
	table[9] = []string{"w", "x", "y", "z"}

	var res []string
	var dfs func(int, string)

	dfs = func(at int, path string) {
		if len(path) == len(digits) {
			res = append(res, path)
			return
		}

		if !(at < len(digits)) { return }

		digit, _ := strconv.Atoi(string(digits[at]))
		comb := table[digit]

		for i:=0; i<len(comb); i++ {
			dfs(at+1, path+comb[i])
		}
	}

	dfs(0, "")

	return res
}
