package main

import "testing"

func TestAllInOneTable(t *testing.T) {
	ts := []struct {
		Op      Operation
		A       int
		B       int
		WantRes int
		WantErr error
	}{
		{Op: OpAdd, A: 1, B: 2, WantRes: 3, WantErr: nil},
		{Op: OpSub, A: 5, B: 3, WantRes: 2, WantErr: nil},
		{Op: OpMul, A: 5, B: 3, WantRes: 15, WantErr: nil},
		{Op: OpDiv, A: 15, B: 3, WantRes: 5, WantErr: nil},
		{Op: OpDiv, A: 15, B: 0, WantRes: 0, WantErr: ErrInvalidB},
	}

	c := Calculator{}

	for _, tc := range ts {
		a, b := tc.A, tc.B
		op := tc.Op
		want := tc.WantRes

		res, err := c.Do(op, a, b)
		if err != tc.WantErr {
			t.Errorf("do error: %v", err)
			continue
		}

		if res != want {
			t.Errorf("calculater error: a=%d,b=%d,res=%d,want=%d", a, b, res, want)
		}
	}
}
