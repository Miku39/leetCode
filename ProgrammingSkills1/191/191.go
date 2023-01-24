package main

// 191. Number of 1 Bits
func hammingWeight(num uint32) int {
	result := 0
	for {
		if num%2 == 1 {
			result++
		}
		num = num / 2
		if num == 0 {
			break
		}
	}
	return result
}
