package main

import (
	"strconv"
)

// 1309. Decrypt String from Alphabet to Integer Mapping
func freqAlphabets(s string) string {
	ASCI_INT := 96
	result := ""
	changeNumber := 0
	for i := 0; i < len(s); i++ {
		if i+2 < len(s) && s[i+2] == byte('#') {
			changeNumber, _ = strconv.Atoi(string(s[i]) + string(s[i+1]))
			i += 2
		} else {
			changeNumber, _ = strconv.Atoi(string(s[i]))
		}
		result += string(rune(ASCI_INT + changeNumber))
	}
	return result
}
