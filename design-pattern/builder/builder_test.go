package main

import (
	"strings"
	"testing"
)

func TestBuildJSON(t *testing.T) {
	var result string
	{
		b := newJSONBuilder()

		d := director{b}
		d.build()

		result = b.getResult()
	}

	if !strings.HasPrefix(result, "{") {
		t.Errorf("json should has prefix with {")
	}

	if !strings.HasSuffix(result, "}") {
		t.Errorf("json should has suffix with }")
	}
}
