package main

// 1431. Kids With the Greatest Number of Candies
// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/?envType=study-plan-v2&envId=leetcode-75
func kidsWithCandies(candies []int, extraCandies int) []bool {
	greatest := 0
	for _, v := range candies {
		if greatest < v {
			greatest = v
		}
	}

	result := make([]bool, len(candies))
	for i, v := range candies {
		if v+extraCandies >= greatest {
			result[i] = true
		}
	}

	return result
}
