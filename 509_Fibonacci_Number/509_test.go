package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	got := fib(2)
	want := 1

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestFib2(t *testing.T) {
	got := fib(3)
	want := 2

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestFib3(t *testing.T) {
	got := fib(4)
	want := 3

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
