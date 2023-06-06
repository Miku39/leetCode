package main

import "testing"

/*
*
Input
["ParkingSystem", "addCar", "addCar", "addCar", "addCar"]
[[1, 1, 0], [1], [2], [3], [1]]
Output
[null, true, true, false, false]

Explanation
ParkingSystem parkingSystem = new ParkingSystem(1, 1, 0);
parkingSystem.addCar(1); // return true because there is 1 available slot for a big car
parkingSystem.addCar(2); // return true because there is 1 available slot for a medium car
parkingSystem.addCar(3); // return false because there is no available slot for a small car
parkingSystem.addCar(1); // return false because there is no available slot for a big car. It is already occupied.
*
*/
func TestParkingSystem(t *testing.T) {
	parkingSystem := Constructor(1, 1, 0)
	got1 := parkingSystem.AddCar(1)
	if got1 != true {
		t.Errorf("got %v, wanted %v", got1, true)
	}
	got2 := parkingSystem.AddCar(2)
	if got2 != true {
		t.Errorf("got %v, wanted %v", got2, true)
	}
	got3 := parkingSystem.AddCar(3)
	if got3 != false {
		t.Errorf("got %v, wanted %v", got3, false)
	}
	got4 := parkingSystem.AddCar(1)
	if got4 != false {
		t.Errorf("got %v, wanted %v", got4, false)
	}
}
