package main

// 283. Move Zeroes
func moveZeroes(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	for i, num := range nums {
		if num != 0 {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if nums[j] != 0 {
				nums[i] = nums[j]
				nums[j] = 0
				break
			}
		}
	}
	return nums
}
