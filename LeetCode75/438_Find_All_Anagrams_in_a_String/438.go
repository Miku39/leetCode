package main

// 438. Find All Anagrams in a String
// https://leetcode.com/problems/find-all-anagrams-in-a-string/
func findAnagrams(s string, p string) []int {
	pMap := make(map[byte]int)
	sMap := make(map[byte]int)

	for i := 0; i < len(p); i++ {
		pMap[p[i]]++
	}

	result := []int{}
	for i := 0; i < len(s); i++ {
		if i-len(p) >= 0 {
			sMap[s[i-len(p)]]--

			if sMap[s[i-len(p)]] == 0 {
				delete(sMap, s[i-len(p)])
			}
		}

		sMap[s[i]]++

		if len(pMap) == len(sMap) {
			isSame := true

			for key, val := range pMap {
				if v, ok := sMap[key]; !ok {
					isSame = false
					break
				} else {
					if v != val {
						isSame = false
						break
					}
				}
			}

			if isSame {
				result = append(result, i+1-len(p))
			}
		}
	}

	return result
}
