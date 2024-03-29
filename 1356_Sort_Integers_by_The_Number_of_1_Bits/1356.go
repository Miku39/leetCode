package main

import (
	"math/bits"
	"sort"
)

// 1356. Sort Integers by The Number of 1 Bits
// https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/
func sortByBits(arr []int) []int {
	sort.Ints(arr)
	bitMap := make(map[int][]int)
	mapKey := make([]int, 0, 500)
	for _, v := range arr {
		count := bits.OnesCount(uint(v))
		array, ok := bitMap[count]
		if !ok {
			array = make([]int, 0, 500)
			mapKey = append(mapKey, count)
		}
		array = append(array, v)
		bitMap[count] = array
	}

	sort.Ints(mapKey)
	result := make([]int, 0, 500)
	for _, x := range mapKey {
		result = append(result, bitMap[x]...)
	}
	return result
}
