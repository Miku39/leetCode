package main

// 643. Maximum Average Subarray I
func findMaxAverage(nums []int, k int) float64 {
	previousSum := 0
	count := 0
	maxSum := 0
	for i, v := range nums {
		if count < k {
			previousSum += v
			count++
			maxSum = previousSum
			continue
		}

		previousSum = previousSum - nums[i-k] + v
		if maxSum < previousSum {
			maxSum = previousSum
		}
	}

	return float64(maxSum) / float64(k)
}
