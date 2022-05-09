package general

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string {}
	}

	numberToLettersMap := make(map[rune][]string)
	numberToLettersMap['2'] = []string { "a", "b", "c" }
	numberToLettersMap['3'] = []string { "d", "e", "f" }
	numberToLettersMap['4'] = []string { "g", "h", "i" }
	numberToLettersMap['5'] = []string { "j", "k", "l" }
	numberToLettersMap['6'] = []string { "m", "n", "o" }
	numberToLettersMap['7'] = []string { "p", "q", "r", "s" }
	numberToLettersMap['8'] = []string { "t", "u", "v" }
	numberToLettersMap['9'] = []string { "w", "x", "y", "z" }

	letterCombinations := make([]string, 0, 0)
	letterCombinations = append(letterCombinations, numberToLettersMap[rune(digits[0])]...)

	for index := 1; index < len(digits); index++ {
		numberLetters := numberToLettersMap[rune(digits[index])]

		lettersToProcess := letterCombinations[:]
		letterCombinations = make([]string, 0, 0)

		for _, letter := range lettersToProcess {
			for _, numberLetter := range numberLetters {
				letterCombinations = append(letterCombinations, letter + numberLetter)
			}
		}
	}

	return letterCombinations
}
