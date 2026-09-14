func isAnagram(s string, t string) bool {
	if len(s) != len(t) { return false }

	var anas [26]int
	var anat [26]int

	for i, val := range s {
		anas[val - 'a']++
		anat[t[i] - 'a']++
	}

	for i, val := range anas {
		if val != anat[i] { return false }
	}

	return true
}
