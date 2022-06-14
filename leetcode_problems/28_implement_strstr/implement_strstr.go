package _28_implement_strstr

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	if len(needle) == 0 {
		return 0
	}

	scanStart := 0
	scanIndex := 0

	for index := 0; index < len(haystack); index++ {
		if haystack[index] == needle[scanIndex] {
			scanIndex++

			if scanIndex == len(needle) {
				return scanStart
			}
		} else {
			index = scanStart
			scanStart = index + 1
			scanIndex = 0
		}
	}

	return -1
}

