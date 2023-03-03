package main

import (
	"testing"
)

func TestGetHint(t *testing.T) {
	got := getHint("1807", "7810")
	want := "1A3B"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
