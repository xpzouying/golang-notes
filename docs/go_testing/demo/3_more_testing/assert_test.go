package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiv(t *testing.T) {
	res, err := div(8, 4)
	want := 2
	assert.NoError(t, err)
	assert.Equal(t, want, res)
}
