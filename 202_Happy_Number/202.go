package main

// 202. Happy Number
func isHappy(n int) bool {
	previous := []int{n}

	for {
		n = calc(n)
		if n == 1 {
			return true
		}
		if isDuplicate(n, previous) {
			return false
		}
		previous = append(previous, n)
	}
}

func calc(num int) int {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num = num / 10
	}
	return sum
}

func isDuplicate(num int, previous []int) bool {
	for _, v := range previous {
		if v == num {
			return true
		}
	}
	return false
}
