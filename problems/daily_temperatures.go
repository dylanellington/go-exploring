package problems

type Temperature struct {
	index int
	value int
}

func dailyTemperatures(temperatures []int) []int {
	stack := make([]Temperature, 0)
	warmDays := make([]int, len(temperatures))

	for index, temp := range temperatures {
		// empty stack
		if len(stack) == 0 {
			stack = append(stack, Temperature { index: index, value: temp, })
			continue
		}

		pastTemperature := stack[len(stack) - 1]

		// current day is a lower temperature than the last
		if pastTemperature.value > temp {
			stack = append(stack, Temperature { index: index, value: temp, })
			continue
		}

		// current day is a higher temperature, process any lower temperature days that came before it
		for pastTemperature.value < temp {
			dayDistance := index - pastTemperature.index
			warmDays[pastTemperature.index] = dayDistance
			stack = stack[:len(stack) - 1]

			if len(stack) == 0 {
				break
			}

			pastTemperature = stack[len(stack) - 1]
		}

		stack = append(stack, Temperature { index: index, value: temp, })
	}

	return warmDays
}
