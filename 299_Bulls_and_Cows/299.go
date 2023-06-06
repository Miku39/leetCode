package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 299. Bulls and Cows
// https://leetcode.com/problems/bulls-and-cows/
func getHint(secret string, guess string) string {
	arraySecret := strings.Split(secret, "")
	arrayGuess := strings.Split(guess, "")

	arrayInt := make([]int, 10)
	countBull := 0
	countCow := 0
	for i := range arraySecret {
		if arraySecret[i] == arrayGuess[i] {
			countBull++
			continue
		} else {
			numberSecret, _ := strconv.Atoi(arraySecret[i])
			numberGuess, _ := strconv.Atoi(arrayGuess[i])
			if arrayInt[numberSecret] < 0 {
				countCow++
			}
			if arrayInt[numberGuess] > 0 {
				countCow++
			}
			arrayInt[numberSecret]++
			arrayInt[numberGuess]--
		}
	}

	return fmt.Sprintf("%dA%dB", countBull, countCow)
}
