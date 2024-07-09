package main

func removeDuplicates(nums []int) int {
	seem := make(map[int]bool)
	count := 0

	arrLen := len(nums) - 1
	for i := 0; i < arrLen; i++ {
		if _, ok := seem[nums[i]]; ok {
			for nums[i+1] == nums[i] {
				nums = append(nums[:i], nums[i+1:]...)
				i--
				arrLen--
			}
		}

		// len = 10
		// i = 1
		// 0 0 0 0 1 1 2 2 3 3
		count++
		seem[nums[i]] = true
	}

	return 0
}
