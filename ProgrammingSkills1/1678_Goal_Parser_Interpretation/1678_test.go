package main

import "testing"

func TestInterpret(t *testing.T) {
	got := interpret("G()(al)")
	want := "Goal"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestInterpret2(t *testing.T) {
	got := interpret("G()()()()(al)")
	want := "Gooooal"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestInterpret3(t *testing.T) {
	got := interpret("(al)G(al)()()G")
	want := "alGalooG"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
