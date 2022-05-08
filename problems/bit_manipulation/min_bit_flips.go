package problems

func minBitFlips(start int, goal int) int {
	xorResult := start^goal
	bitFlips := 0

	for xorResult > 0 {
		isOne := xorResult % 2 != 0

		if isOne {
			bitFlips++
		}

		xorResult >>= 1
	}

	return bitFlips
}