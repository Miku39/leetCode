package main

// 1603. Design Parking System
// https://leetcode.com/problems/design-parking-system/
type ParkingSystem struct {
	carType map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		carType: map[int]int{
			1: big,
			2: medium,
			3: small,
		},
	}
}

func (p *ParkingSystem) AddCar(carType int) bool {
	if p.carType[carType] > 0 {
		p.carType[carType]--
		return true
	}

	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
