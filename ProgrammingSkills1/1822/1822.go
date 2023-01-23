package main

import "fmt"

func main() {
	input := []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}
	result := arraySign(input)
	fmt.Println(result)
}

// 1822. Sign of the Product of an Array
func arraySign(nums []int) int {
	countMinus := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}

		if num < 0 {
			countMinus++
		}
	}

	if countMinus%2 == 0 {
		return 1
	}

	return -1
}
