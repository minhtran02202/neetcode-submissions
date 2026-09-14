func isAnagram(s string, t string) bool {
	if len(s) != len(t) { return false }

	var freq [26]int

	for i := range s { freq[s[i] - 'a']++; freq[t[i] - 'a']-- }

	for _, val := range freq { if val != 0 { return false } }

	return true
}
