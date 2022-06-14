package _4_median_of_two_sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArray := mergeSortedArrays(nums1, nums2)

	if len(mergedArray) % 2 == 0 {
		index := len(mergedArray) / 2
		return (float64(mergedArray[index]) + float64(mergedArray[index - 1])) / 2
	}

	index := len(mergedArray) / 2
	return float64(mergedArray[index])
}

func mergeSortedArrays(arrayOne []int, arrayTwo []int) []int {
	// skip merge process if either array is empty
	if len(arrayOne) == 0 {
		return arrayTwo
	}

	if len(arrayTwo) == 0 {
		return arrayOne
	}

	// allocate exact size for merged array
	mergedArraySize := len(arrayOne) + len(arrayTwo)
	mergedArray := make([]int, 0, mergedArraySize)

	arrayOneLength, arrayOneIndex := len(arrayOne), 0
	arrayTwoLength, arrayTwoIndex := len(arrayTwo), 0

	// merge arrays until at least one array is completely copied
	for arrayOneIndex < arrayOneLength && arrayTwoIndex < arrayTwoLength {
		if arrayOne[arrayOneIndex] < arrayTwo[arrayTwoIndex] {
			mergedArray = append(mergedArray, arrayOne[arrayOneIndex])
			arrayOneIndex++
		} else {
			mergedArray = append(mergedArray, arrayTwo[arrayTwoIndex])
			arrayTwoIndex++
		}
	}

	// merge the remaining elements of the other array
	for arrayOneIndex < arrayOneLength {
		mergedArray = append(mergedArray, arrayOne[arrayOneIndex])
		arrayOneIndex++
	}

	for arrayTwoIndex < arrayTwoLength {
		mergedArray = append(mergedArray, arrayTwo[arrayTwoIndex])
		arrayTwoIndex++
	}

	return mergedArray
}
