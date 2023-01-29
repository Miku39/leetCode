package main

import "testing"

func TestFreqAlphabets(t *testing.T) {
	got := freqAlphabets("10#11#12")
	want := "jkab"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestFreqAlphabets2(t *testing.T) {
	got := freqAlphabets("1326#")
	want := "acz"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
