package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{3, 4, 15, 2, 9, 4}
	result := largestPerimeter(input)
	fmt.Println(result)
}

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i] < (nums[i-1] + nums[i-2]) {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}
