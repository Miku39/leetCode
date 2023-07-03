package main

// 283. Move Zeroes
func moveZeroes(nums []int) []int {
	nonZeroIndex := 0
	for i, num := range nums {
		if num != 0 {
			nums[i], nums[nonZeroIndex] = nums[nonZeroIndex], nums[i]
			nonZeroIndex++
		}
	}

	return nums
}
