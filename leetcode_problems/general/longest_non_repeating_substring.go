package general

func LongestNonRepeatingSubstring(s string) int {
	hashmap := make(map[rune]int)
	substringIndex := 0
	longestCount := 0

	for index, character := range s {
		savedIndex, exists := hashmap[character]

		// move start of substring marker
		if exists && savedIndex >= substringIndex {
			substringIndex = savedIndex + 1
		}

		// calculate substring length and compare
		length := (index + 1) - substringIndex

		if length > longestCount {
			longestCount = length
		}

		// save current character
		hashmap[character] = index
	}

	return longestCount
}
