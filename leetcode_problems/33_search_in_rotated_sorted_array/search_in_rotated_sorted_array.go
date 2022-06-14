package _33_search_in_rotated_sorted_array

func search(nums []int, target int) int {
	targetIndex := 0

	// binary search through the array of numbers
	for len(nums) > 1 {
		middleIndex := len(nums) / 2
		lastIndex := len(nums) - 1

		// left half is continuous
		if nums[0] < nums[middleIndex] {
			// left half could contain the target, so search it
			if target >= nums[0] && target < nums[middleIndex] {
				nums = nums[0:middleIndex]
				continue
			}

			// search the right half instead
			nums = nums[middleIndex:len(nums)]
			targetIndex += middleIndex
			// right half is continuous
		} else {
			// right half could contain the target, so search it
			if target >= nums[middleIndex] && target <= nums[lastIndex] {
				nums = nums[middleIndex:len(nums)]
				targetIndex += middleIndex
				continue
			}

			// search the left half instead
			nums = nums[0:middleIndex]
		}
	}

	if nums[0] != target {
		return -1
	}

	return targetIndex
}
