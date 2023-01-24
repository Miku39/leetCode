package main

import (
	"sort"
)

// 1502. Can Make Arithmetic Progression From Sequence
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	baseDifference := arr[1] - arr[0]
	for i := len(arr) - 1; i >= 1; i-- {
		if (arr[i] - arr[i-1]) != baseDifference {
			return false
		}
	}

	return true
}
