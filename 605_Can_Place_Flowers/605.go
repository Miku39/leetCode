package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i, v := range flowerbed {
		if v == 1 {
			continue
		}

		isPreviousEmpty := i == 0 || flowerbed[i-1] == 0
		isNextEmpty := i == (len(flowerbed)-1) || flowerbed[i+1] == 0

		if isPreviousEmpty && isNextEmpty {
			count++
			flowerbed[i] = 1
		}
	}

	return count >= n
}
