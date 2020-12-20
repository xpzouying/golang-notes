package main

import "testing"

func TestMod(t *testing.T) {

	got, err := Mod(5, 3)
	if err != nil {
		t.Error(err)
	}
	want := 2
	if got != want {
		t.Errorf("expected error")
	}

	got2, err := Mod(5, 2)
	if err != nil {
		t.Error(err)
	}
	want2 := 1
	if got2 != want2 {
		t.Error("expected error")
	}
}
