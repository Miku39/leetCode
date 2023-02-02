package main

// 1480. Running Sum of 1d Array
func runningSum(nums []int) []int {
	var result []int
	var sum int
	for _, num := range nums {
		sum += num
		result = append(result, sum)
	}
	return result
}
