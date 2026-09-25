func isSubsequence(s string, t string) bool {
	if len(s) > len(t) { return false }

	ps, pt := 0, 0

	for ps < len(s) && pt < len(t) {
		if s[ps] == t[pt] { ps++ }
		pt ++
	}

	return ps == len(s)
}
