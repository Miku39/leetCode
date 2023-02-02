package main

// 389. Find the Difference
func findTheDifference(s string, t string) byte {
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}

	for j := 0; j < len(t); j++ {
		tMap[t[j]]++
	}

	for key, tCount := range tMap {
		sCount, ok := sMap[key]
		if !ok || tCount != sCount {
			return key
		}
	}

	return s[0]
}
