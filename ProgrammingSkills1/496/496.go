package main

// 496. Next Greater Element I
func nextGreaterElement(nums1 []int, nums2 []int) []int {

	result := make([]int, len(nums1))

	for i, v := range nums1 {

		flag := false
		for _, w := range nums2 {
			if v == w {
				flag = true
			}

			if !flag {
				continue
			}

			if v < w {
				result[i] = w
				break
			}
		}

		if result[i] == 0 {
			result[i] = -1
		}
	}

	return result
}
