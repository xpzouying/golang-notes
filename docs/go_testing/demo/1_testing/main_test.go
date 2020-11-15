package main

import "testing"

func TestAdd(t *testing.T) {
	c := Calculator{}

	a, b := 1, 2
	op := OpAdd

	res, err := c.Do(op, a, b)
	if err != nil {
		t.Errorf("do error: %v", err)
	}

	want := 3
	if res != want {
		t.Errorf("calculater error: a=%d,b=%d,res=%d,want=%d", a, b, res, want)
	}
}

func TestDivNormal(t *testing.T) {
	c := Calculator{}

	a, b := 10, 2
	op := OpDiv

	res, err := c.Do(op, a, b)
	if err != nil {
		t.Errorf("do error: %v", err)
	}

	want := 5
	if res != want {
		t.Errorf("calculater error: a=%d,b=%d,res=%d,want=%d", a, b, res, want)
	}
}

func TestDivZero(t *testing.T) {
	c := Calculator{}

	a, b := 10, 0
	op := OpDiv

	_, err := c.Do(op, a, b)
	if err == nil {
		t.Errorf("div 0 should failed")
	}
}
