func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {return false}

	h1 := make([]int, 26)
	h2 := make([]int, 26)

    for i := 0; i < len(s1); i++ {
        h1[s1[i]-'a']++
        h2[s2[i]-'a']++
    }

	s1L := len(s1)
	for i:=s1L; i < len(s2); i++ {
		if match(h1, h2) {
			return true
		}
		h2[s2[i - s1L] - 'a']--
		h2[s2[i] - 'a']++
	}

	return match(h1, h2)
}

func match(h1, h2 []int) bool {
	for i:=0; i<26; i++ {
		if h1[i] != h2[i] {
			return false
		}
	}

	return true
}
