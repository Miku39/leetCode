package main

// 217. Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	countMap := make(map[int]int)

	for _, v := range nums {
		count := countMap[v]
		if count == 1 {
			return true
		}

		countMap[v] = 1
	}

	return false
}
