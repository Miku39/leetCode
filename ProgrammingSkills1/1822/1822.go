package main

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
