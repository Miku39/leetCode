package main

// 1588. Sum of All Odd Length Subarrays
func sumOddLengthSubarrays(arr []int) int {
	result := 0

	for i := 1; i <= len(arr); i++ {
		if i%2 == 0 {
			continue
		}

		for k := 0; k < i; k++ {
			for j := 0; j < len(arr); j++ {
				if j+i > len(arr) {
					break
				}
				result += arr[j+k]
			}
		}
	}

	return result
}
