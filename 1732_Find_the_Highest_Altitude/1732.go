package main

// 1732. Find the Highest Altitude
func largestAltitude(gain []int) int {
	altitude := 0
	highestAltitude := 0
	for _, v := range gain {
		altitude += v
		if highestAltitude < altitude {
			highestAltitude = altitude
		}
	}

	return highestAltitude
}
