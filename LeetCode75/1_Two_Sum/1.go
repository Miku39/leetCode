package main

// 1. Two Sum
// https://leetcode.com/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // num, index
	for index, num := range nums {
		mapIndex, ok := numMap[target-num]
		if ok {
			return []int{mapIndex, index}
		}
		numMap[num] = index
	}

	return nil
}
