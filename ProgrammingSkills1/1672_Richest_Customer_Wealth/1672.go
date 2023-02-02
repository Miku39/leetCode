package main

// 1672. Richest Customer Wealth
func maximumWealth(accounts [][]int) int {
	maxWealth := 0

	for _, account := range accounts {

		wealth := 0
		for _, w := range account {
			wealth += w
		}

		if wealth > maxWealth {
			maxWealth = wealth
		}
	}

	return maxWealth
}
