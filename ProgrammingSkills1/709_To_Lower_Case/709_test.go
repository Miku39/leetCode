package main

import "testing"

func TestToLowerCase(t *testing.T) {
	got := toLowerCase("Hello")
	want := "hello"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
