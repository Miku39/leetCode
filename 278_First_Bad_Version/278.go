package main

// 278. First Bad Version
// https://leetcode.com/problems/first-bad-version/?envType=study-plan&id=level-1

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool {
	return version >= 4
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	mid := 1
	for left < right {
		mid = (left + right) / 2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
