package main

import "fmt"

func main() {
	result := subtractProductAndSum(4421)
	fmt.Println(result)
}

// 1281. Subtract the Product and Sum of Digits of an Integer
func subtractProductAndSum(n int) int {
	product := 1
	sum := 0
	for n != 0 {
		num := n % 10
		product *= num
		sum += num
		n = n / 10
	}
	return product - sum
}
