package main

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	flowerbed := []int{1, 0, 0, 0, 1}
	got := canPlaceFlowers(flowerbed, 1)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCanPlaceFlowers2(t *testing.T) {
	flowerbed := []int{1, 0, 0, 0, 1}
	got := canPlaceFlowers(flowerbed, 2)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
