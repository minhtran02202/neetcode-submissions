func longestPalindrome(s string) string {
    n := len(s)
	if n==0 {
		return s
	}

	memo := make(map[[2]int]string)

	var find func(int, int) string
	find = func(start, end int) string {
		if start > end { return "" }
		if start == end { return string(s[start]) }

		state := [2]int{start, end}

		if val, ok := memo[state]; ok {
			return val
		}

		if s[start] == s[end] {
			inner := find(start + 1, end - 1)
			if len(inner) == end - start - 1 {
				memo[state] = s[start : end+1]
				return memo[state]
			}
		} 

		l := find(start + 1, end)
		r := find(start, end - 1)

		res := r
		if len(l) > len(r) {
			res = l
		}

		memo[state] = res
		return res
	}

	return find(0, len(s) - 1)
}
