func lengthOfLongestSubstring(s string) int {
	res := 0
	hash := make(map[byte]bool)
	l := 0

	for r:=0; r<len(s); r++ {
		if !hash[s[r]] {
			hash[s[r]] = true
		} else {
			for l < r {
				if s[l] == s[r] {
					l++
					break
				}
				hash[s[l]] = false
				l++
			}
		}

		localMax := r - l + 1
		res = max(res, localMax)
	}

	return res
}
