package main

import "testing"

func TestAreAlmostEqual(t *testing.T) {

	got := areAlmostEqual("abcd", "dcba")
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestAreAlmostEqual2(t *testing.T) {

	got := areAlmostEqual(
		"eyapgosevqecyuobikhcpztfrrnuibtpjekygfssnqihtmlnevrwqosktoozmefq",
		"iexryyjgmazbxegpxxipzoybexebydjwxgarqiysoqexgvelhlbpotlfbnogbvpg",
	)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
