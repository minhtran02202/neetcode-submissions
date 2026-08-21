func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    hashS := make(map[byte]int)
    hashT := make(map[byte]int)

    for i := 0; i < len(s); i++ {
        hashS[s[i]]++
        hashT[t[i]]++
    }

    for k := range hashS {
        if hashT[k] != hashS[k] {
            return false
        }
    }

    // If there are keys only in hashT but not in hashS, the previous loop won't catch them.
    // You can either check the sizes or iterate hashT too:
    for k := range hashT {
        if hashS[k] != hashT[k] {
            return false
        }
    }

    return true
}