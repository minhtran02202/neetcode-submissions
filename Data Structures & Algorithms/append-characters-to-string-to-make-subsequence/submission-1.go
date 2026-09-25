// find i index of t that is last valid subsequent char
// return len t - i (since once valid, i auto move to next char, no need to handle offset by 1)
func appendCharacters(s string, t string) int {
    ps, pt := 0, 0

	for ps < len(s) && pt < len(t) {
		if s[ps] == t[pt] { pt++ }
		ps++
	}

	return len(t) - pt
}