func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hits := map[byte]int{}
	for i:= 0; i < len(s); i++ {
		sByte, tByte := []byte(s)[i], []byte(t)[i]
		hits[sByte]++
		hits[tByte]--
	}

	for _, v := range hits {
		if v!=0 {
			return false
		}
	}
	return true
}
