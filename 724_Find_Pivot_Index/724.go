package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}

// 724. Find Pivot Index
func pivotIndex(nums []int) int {
	result := -1
	for j := 0; j < len(nums); j++ {
		var leftSum int
		var rightSum int
		for i, num := range nums {
			switch {
			case i < j:
				leftSum += num
			case i > j:
				rightSum += num
			default:
				continue
			}
		}
		if leftSum == rightSum {
			result = j
			break
		}
	}
	return result
}
