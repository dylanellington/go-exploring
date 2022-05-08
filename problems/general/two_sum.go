package problems

func twoSum(nums []int, target int) []int {
	for j := 0; j < len(nums) - 1; j++ {
		for k := j + 1; k < len(nums); k++ {
			if nums[j] + nums[k] == target {
				return []int{j, k}
			}
		}
	}

	return []int{}
}

func twoSumOptimal(nums []int, target int) []int {
	dict := make(map[int]int)
	dict[nums[0]] = 0

	for index := 1; index < len(nums); index++ {
		sumIndex, exists := dict[target - nums[index]]

		if exists {
			return []int{sumIndex, index}
		}

		dict[nums[index]] = index
	}

	return []int{}
}
