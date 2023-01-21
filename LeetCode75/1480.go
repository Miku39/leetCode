package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
}

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
