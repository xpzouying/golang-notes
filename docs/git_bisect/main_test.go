package main

import "testing"

func TestMod(t *testing.T) {

	got := Mod(5, 3)
	want := 2
	if got != want {
		t.Errorf("expected error")
	}

	got2 := Mod(5, 2)
	want2 := 1
	if got2 != want2 {
		t.Error("expected error")
	}
}
