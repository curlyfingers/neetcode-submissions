import "slices"
func groupAnagrams(strs []string) [][]string {
	patterns :=  map[string][]string{}
	for _, str :=  range strs {
		letters := []byte{}
		letterFreq := map[byte]int{}
		for _, b := range []byte(str) {
			if _, ok := letterFreq[b]; !ok {
				letters = append(letters, b)
			}
			letterFreq[b]++
		}
		slices.Sort(letters)

		var key string
		for _, letter := range letters {
			key = fmt.Sprintf("%s%s%d", key, string([]byte{letter}), letterFreq[letter])
		}
		if _, ok := patterns[key]; !ok {
			patterns[key] = []string{str}
			continue
		}
		patterns[key] = append(patterns[key], str)
	}

	results := make([][]string, 0, len(patterns))
	for _, group := range patterns {
		results = append(results, group)
	}
	return results
}
