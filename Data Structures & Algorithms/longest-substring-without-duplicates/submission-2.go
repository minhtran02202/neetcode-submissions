func lengthOfLongestSubstring(s string) int {
	l, res := 0, 0
	mp := make(map[rune]int)

	for r, val := range s {
		if i, ok := mp[val]; ok {
			l = max(i+1, l)
		}
		mp[val] = r
		res = max(res, r-l+1)
	}

	return res
}
