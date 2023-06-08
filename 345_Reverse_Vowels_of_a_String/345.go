package main

// 345. Reverse Vowels of a String
// https://leetcode.com/problems/reverse-vowels-of-a-string/?envType=study-plan-v2&envId=leetcode-75
func reverseVowels(s string) string {
	vowels := []byte{}
	for i := 0; i < len(s); i++ {
		if isVowels(s[i]) {
			vowels = append(vowels, s[i])
		}
	}

	vowelsIndex := len(vowels) - 1
	result := ""
	for i := 0; i < len(s); i++ {
		addLetter := s[i]
		if isVowels(s[i]) {
			addLetter = vowels[vowelsIndex]
			vowelsIndex--
		}
		result += string(addLetter)
	}

	return result
}

func isVowels(b byte) bool {
	if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
		b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U' {
		return true
	}
	return false
}
