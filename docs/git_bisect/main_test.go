package main

import "testing"

func TestMod(t *testing.T) {

	got := Mod(5, 3)
	want := 2
	if got != want {
		t.Errorf("expected error")
	}
}
