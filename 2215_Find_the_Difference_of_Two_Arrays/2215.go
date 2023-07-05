package main

// 2215. Find the Difference of Two Arrays
func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1Map := make(map[int]bool)
	for _, v := range nums1 {
		nums1Map[v] = true
	}

	nums2Map := make(map[int]bool)
	for _, v := range nums2 {
		nums2Map[v] = true
	}

	result := make([][]int, 2)
	for _, v := range nums1 {
		_, ok := nums2Map[v]
		if !ok {
			result[0] = append(result[0], v)
			nums2Map[v] = false
		}
	}

	for _, v := range nums2 {
		_, ok := nums1Map[v]
		if !ok {
			result[1] = append(result[1], v)
			nums1Map[v] = false
		}
	}

	return result
}
