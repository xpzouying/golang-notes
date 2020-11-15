package main

import "testing"

func TestSkipTestcase(t *testing.T) {
	t.Skip("skip this testcase")

	t.Logf("at the end of function")
}
