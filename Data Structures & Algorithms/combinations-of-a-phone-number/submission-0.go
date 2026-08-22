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
	var path string
	var dfs func(int)

	dfs = func(at int) {
		if len(path) == len(digits) {
			res = append(res, path)
			return
		}

		if !(at < len(digits)) { return }

		digit, _ := strconv.Atoi(string(digits[at]))
		comb := table[digit]

		for i:=0; i<len(comb); i++ {
			path += comb[i]
			dfs(at+1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)

	return res
}

// 	table[2] = "abc"
	// table[3] = "def"
	// table[4] = "ghi"
	// table[5] = "jkl"
	// table[6] = "mno"
	// table[7] = "pqrs"
	// table[8] = "tuv"
	// table[9] = "wxyz"