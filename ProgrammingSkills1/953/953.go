package main

// 953. Verifying an Alien Dictionary
func isAlienSorted(words []string, order string) bool {
	orderMap := make(map[byte]int)
	for i, _ := range order {
		orderMap[order[i]] = i
	}

	for j := 0; j < len(words)-1; j++ {
		for k, _ := range words[j] {
			if k == len(words[j+1]) {
				return false
			}

			if orderMap[words[j][k]] < orderMap[words[j+1][k]] {
				break
			} else if orderMap[words[j][k]] > orderMap[words[j+1][k]] {
				return false
			}
		}
	}

	return true
}
